<template>
  <div class="stores-container">
    <!-- Loading & Error States -->
    <div v-if="loading" class="loading-state" role="status">Loading stores...</div>
    <div v-else-if="error" class="error-state" role="alert">
      <p>Error: {{ error }}</p>
      <button @click="retryFetchStores()" class="button-base button-primary" :disabled="processing">
        Retry
      </button>
    </div>

    <!-- Header with Search Bar and Add Button -->
    <div class="page-header">
      <input
        type="text"
        v-model="searchTerm"
        placeholder="Search stores..."
        class="search-bar"
        aria-label="Search stores"
      />
      <button @click="openAddStoreModal" class="button-base button-primary" :disabled="processing">
        Add Store
      </button>
    </div>

    <!-- Main Content -->
    <div class="content-wrapper">
      <!-- Stores Grid -->
      <div class="stores-grid" role="grid">
        <div
          v-for="store in filteredStores"
          :key="store.store_id"
          class="store-card"
          @click="showStoreDetails(store)"
          role="gridcell"
        >
          <div class="store-content">
            <h3 class="store-title">{{ store.store_name }}</h3>
            <div class="store-details">
              <p><strong>Branch:</strong> {{ store.store_branch }}</p>
              <p><strong>Address:</strong> {{ store.store_address }}</p>
              <p><strong>Phone:</strong> {{ store.store_phone }}</p>
            </div>
            <div class="store-actions">
              <button
                @click.stop="editStore(store)"
                class="button-base button-secondary"
                :disabled="processing"
              >
                Edit
              </button>
              <button
                @click.stop="deleteStore(store.store_id)"
                class="button-base button-danger"
                :disabled="processing"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Modals: StoreModal for add/edit, StoreDetailsModal for details -->
  <StoreModal
    v-if="showModal || showAddStoreModal"
    :store="showModal ? editedStore : newStore"
    :isEdit="showModal"
    :processing="processing"
    @save="showModal ? saveEditedStore() : addNewStore()"
    @close="showModal ? closeModal() : closeAddStoreModal()"
  />
  <StoreDetailsModal
    v-if="showStoreDetailsModal"
    :store="selectedStore"
    @close="closeDetailsModal"
  />
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import '@/components/styles/ButtonStyles.vue'
import StoreModal from './StoreModal.vue'
import StoreDetailsModal from './StoreDetailsModal.vue'
import type { Store } from '@/types/Store.ts'

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api'

// State management
const stores = ref<Store[]>([])
const loading = ref(true)

const processing = ref(false)
const error = ref<string | null>(null)
const showModal = ref(false)
const showAddStoreModal = ref(false)
const showStoreDetailsModal = ref(false)
const selectedStore = ref<Store | null>(null)

// Search term for filtering stores
const searchTerm = ref('')

// Computed property for filtered stores based on the search term
const filteredStores = computed(() => {
  if (!searchTerm.value) return stores.value
  const term = searchTerm.value.toLowerCase()
  return stores.value.filter(
    (store) =>
      store.store_name.toLowerCase().includes(term) ||
      store.store_branch.toLowerCase().includes(term) ||
      store.store_address.toLowerCase().includes(term) ||
      store.store_phone.toLowerCase().includes(term),
  )
})

// Template for edited and new store records
const editedStore = ref<Store>({
  store_id: 0,
  unique_uid: '',
  store_name: '',
  store_branch: '',
  store_address: '',
  store_phone: '',
  updated_at: new Date().toISOString(),
})

const newStore = ref<Store>({
  ...editedStore.value,
})

// API calls

const fetchStores = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await fetch(`${API_BASE}/stores/all`)
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    stores.value = await response.json()
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to fetch stores'
    console.error('Fetch stores error:', e)
  } finally {
    loading.value = false
  }
}

const retryFetchStores = async (retries = 3) => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      await new Promise((resolve) => setTimeout(resolve, attempt * 1000))
      await fetchStores()
      return
    } catch (e) {
      if (attempt === retries) throw e
    }
  }
}

const deleteStore = async (storeId: number) => {
  if (!confirm('Are you sure you want to delete this store?')) return
  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/stores/delete?store_id=${storeId}`, {
      method: 'DELETE',
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    stores.value = stores.value.filter((store) => store.store_id !== storeId)
  } catch (e) {
    error.value = 'Failed to delete store'
    console.error('Delete store error:', e)
  } finally {
    processing.value = false
  }
}

const saveEditedStore = async () => {
  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/stores/update`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(editedStore.value),
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const updatedStore = await response.json()
    stores.value = stores.value.map((store) =>
      store.store_id === editedStore.value.store_id ? updatedStore : store,
    )
    showModal.value = false
  } catch (e) {
    error.value = 'Failed to update store'
    console.error('Update store error:', e)
  } finally {
    processing.value = false
  }
}

const addNewStore = async () => {
  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/stores/create`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newStore.value),
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const createdStore = await response.json()
    stores.value.push(createdStore)
    showAddStoreModal.value = false
  } catch (e) {
    error.value = 'Failed to add store'
    console.error('Add store error:', e)
  } finally {
    processing.value = false
  }
}

// UI Handlers

const editStore = (store: Store) => {
  editedStore.value = { ...store }
  showModal.value = true
}

const showStoreDetails = (store: Store) => {
  selectedStore.value = store
  showStoreDetailsModal.value = true
}

const openAddStoreModal = () => {
  showAddStoreModal.value = true
  newStore.value = {
    store_id: 0,
    unique_uid: '',
    store_name: '',
    store_branch: '',
    store_address: '',
    store_phone: '',
    updated_at: new Date().toISOString(),
  }
}

const closeAddStoreModal = () => {
  showAddStoreModal.value = false
}

const closeModal = () => {
  showModal.value = false
}

const closeDetailsModal = () => {
  showStoreDetailsModal.value = false
  selectedStore.value = null
}

onMounted(() => {
  retryFetchStores()
})
</script>

<style scoped>
.stores-container {
  width: 100%;
  min-height: 100vh;
  margin-top: 64px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  position: absolute;
  left: 0;
  right: 0;
  padding: 0 2rem;
}

.page-header {
  width: 100%;
  max-width: 1800px;
  margin: 0 auto 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  z-index: 1;
}

.search-bar {
  padding: 8px 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-text);
  font-size: 1rem;
  width: 250px;
  transition:
    border-color 0.2s,
    background 0.2s;
}

.search-bar:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.4);
  background: rgba(255, 255, 255, 0.1);
}

.content-wrapper {
  width: 100%;
  max-width: 1800px;
  margin: 0 auto;
}

.stores-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  padding: 20px 0;
  width: 100%;
  justify-content: flex-start;
}

.store-card {
  flex: 0 0 calc(25% - 16px);
  min-width: 280px;
  max-width: calc(25% - 16px);
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.store-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 12px rgba(0, 0, 0, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

.store-content {
  padding: 20px;
}

.store-title {
  margin: 0 0 16px;
  font-size: clamp(1rem, 2vw, 1.2rem);
  color: var(--color-heading);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 12px;
}

.store-details {
  display: grid;
  gap: 12px;
}

.store-details p {
  margin: 0;
  font-size: clamp(0.9rem, 1.8vw, 1rem);
  line-height: 1.4;
}

.store-details strong {
  color: var(--color-heading);
  font-weight: 600;
}

.store-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.loading-state,
.error-state {
  text-align: center;
  padding: 40px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  margin: 20px auto;
  width: 100%;
  max-width: 1200px;
}

.error-state {
  color: #ef4444;
}

@media (max-width: 768px) {
  .stores-container {
    padding: 0 1rem;
  }
  .store-card {
    flex: 0 0 100%;
    max-width: 100%;
  }
  .page-header {
    flex-direction: column;
    align-items: stretch;
    gap: 1rem;
  }
  .search-bar {
    width: 100%;
  }
}
</style>
