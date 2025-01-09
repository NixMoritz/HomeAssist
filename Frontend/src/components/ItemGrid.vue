<template>
  <div class="items-container">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state">Loading items...</div>

    <!-- Error State -->
    <div v-else-if="error" class="error-state">
      <p>Error loading items: {{ error }}</p>
      <button @click="retryFetchItems" class="retry-button">Retry</button>
    </div>

    <!-- Add Item Button -->
    <div class="add-item-container">
      <button @click="openAddItemModal" class="add-item-button">Add Item</button>
    </div>

    <!-- Add Item Modal -->
    <div v-if="showAddItemModal" class="modal-overlay">
      <div class="modal">
        <h2>Add New Item</h2>
        <div class="modal-content">
          <label for="addItemName">Name:</label>
          <input type="text" id="addItemName" v-model="newItem.Item_Name" />

          <label for="addUnitPrice">Price:</label>
          <input type="number" id="addUnitPrice" v-model.number="newItem.Unit_Price" />

          <label for="addUnits">Units:</label>
          <input type="number" id="addUnits" v-model.number="newItem.Units" />

          <label for="addStoreBranch">Store Branch:</label>
          <input type="text" id="addStoreBranch" v-model="newItem.Store_Branch" />

          <label for="addWeight">Weight:</label>
          <input type="number" id="addWeight" v-model.number="newItem.Weight" />
        </div>
        <div class="modal-actions">
          <button @click="addNewItem" class="save-button">Add</button>
          <button @click="closeAddItemModal" class="cancel-button">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Items Grid -->
    <div v-else class="items-grid">
      <div v-for="item in items" :key="item.Item_ID" class="item-card">
        <div class="item-content">
          <h3 class="item-title">{{ item.item_name }}</h3>
          <div class="item-details">
            <p><strong>Price:</strong> {{ formatPrice(item.unit_price) }}</p>
            <p><strong>Units:</strong> {{ item.units }}</p>
            <p><strong>Store Branch:</strong> {{ item.store_branch }}</p>
            <p><strong>Weight:</strong> {{ item.weight }}kg</p>
            <p class="updated-at">Last updated: {{ formatDate(item.updated_at) }}</p>
          </div>
          <div class="item-actions">
            <button @click="editItem(item)" class="edit-button">Edit</button>
            <button @click="deleteItem(item.item_id)" class="delete-button">Delete</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- Edit Modal -->
  <div v-if="showModal" class="modal-overlay">
    <div class="modal">
      <h2>Edit Item</h2>
      <div class="modal-content">
        <label for="editItemName">Name:</label>
        <input type="text" id="editItemName" v-model="editedItem.Item_Name" />

        <label for="editUnitPrice">Price:</label>
        <input type="number" id="editUnitPrice" v-model.number="editedItem.Unit_Price" />

        <label for="editUnits">Units:</label>
        <input type="number" id="editUnits" v-model.number="editedItem.Units" />

        <label for="editStoreBranch">Store Branch:</label>
        <input type="text" id="editStoreBranch" v-model="editedItem.Store_Branch" />

        <label for="editWeight">Weight:</label>
        <input type="number" id="editWeight" v-model.number="editedItem.Weight" />
      </div>
      <div class="modal-actions">
        <button @click="saveEditedItem" class="save-button">Save</button>
        <button @click="closeModal" class="cancel-button">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { toRaw } from 'vue'

interface Item {
  Item_ID: number
  Item_Name: string
  Unit_Price: number
  Units: number
  Store_Branch: string
  Weight: number
  Updated_At: string
}

const items = ref<Item[]>([])
const loading = ref(true)
const error = ref<string | null>(null)
const showModal = ref(false)
const editedItem = ref<Item>({
  Item_ID: 0,
  Item_Name: '',
  Unit_Price: 0,
  Units: 0,
  Store_Branch: '',
  Weight: 0,
  Updated_At: '',
})

const showAddItemModal = ref(false)
const newItem = ref<Item>({
  Item_ID: 0,
  Item_Name: '',
  Unit_Price: 0,
  Units: 0,
  Store_Branch: '',
  Weight: 0,
  Updated_At: new Date().toISOString(),
})

const openAddItemModal = () => {
  showAddItemModal.value = true
  newItem.value = {
    Item_ID: 0,
    Item_Name: '',
    Unit_Price: 0,
    Units: 0,
    Store_Branch: '',
    Weight: 0,
    Updated_At: new Date().toISOString(),
  }
}

const closeAddItemModal = () => {
  showAddItemModal.value = false
}

const addNewItem = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/items/create', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newItem.value),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const createdItem = await response.json()
    items.value.push(createdItem)
    console.log(`New item added successfully:`, createdItem)
    showAddItemModal.value = false
  } catch (error) {
    console.error('Failed to add new item:', error)
  }
}

const fetchItems = async () => {
  try {
    loading.value = true
    error.value = null

    const response = await fetch('http://localhost:8080/api/items', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    })

    if (!response.ok) {
      console.error(`HTTP error! status: ${response.status}`)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const textData = await response.text()
    console.log('Data:', textData) // Log the raw data
    try {
      items.value = JSON.parse(textData)
      console.log('Parsed items:', items.value) // Log the parsed items
    } catch (parseError) {
      console.error('Response was not JSON:', textData)
      throw new Error('Server response was not in JSON format')
    }
  } catch (e) {
    console.error('Fetch error:', e)
    error.value = e instanceof Error ? e.message : 'An unknown error occurred'
  } finally {
    loading.value = false
  }
}

const retryFetchItems = async (retries = 3) => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      await fetchItems()
      return // Exit loop on success
    } catch (e) {
      console.warn(`Attempt ${attempt} failed:`, e)
      if (attempt === retries) throw e
    }
  }
}

onMounted(() => {
  retryFetchItems()
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
  }).format(price)
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('de-DE', {
    weekday: 'short',
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
}

const deleteItem = async (itemId: number) => {
  console.log('Item ID to delete:', itemId)

  if (itemId === undefined || itemId === null) {
    console.error('Invalid item ID')
    return
  }

  if (!confirm('Are you sure you want to delete this item?')) {
    return
  }

  try {
    const response = await fetch(`http://localhost:8080/api/items/delete?item_id=${itemId}`, {
      method: 'DELETE',
    })
    if (!response.ok) {
      console.error(`Failed to delete item, HTTP error! status: ${response.status}`)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    // Use `toRaw` to get the unwrapped item data
    items.value = items.value.filter((item) => toRaw(item).item_id !== itemId)
    console.log(`Item with ID ${itemId} deleted successfully`)
  } catch (error) {
    console.error('Failed to delete item:', error)
  }
}
const editItem = (item: Item) => {
  editedItem.value = { ...item } // Copy item to editedItem for modal
  showModal.value = true
}
const closeModal = () => {
  showModal.value = false
}

const saveEditedItem = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/items/update`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(editedItem.value),
    })

    if (!response.ok) {
      console.error(`Failed to update item, HTTP error! status: ${response.status}`)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const responseData = await response.json()

    items.value = items.value.map((item) =>
      item.Item_ID === editedItem.value.Item_ID ? { ...item, ...editedItem.value } : item,
    )

    console.log(
      `Item with ID ${editedItem.value.Item_ID} updated successfully:`,
      responseData.message,
    )
    showModal.value = false
  } catch (error) {
    console.error('Failed to update item:', error)
  }
}
</script>

<style scoped>
.items-container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.loading-state,
.error-state {
  text-align: center;
  padding: 40px;
  font-size: 1.1rem;
  color: var(--color-text);
}

.error-state p {
  margin-bottom: 1rem;
  color: #dc2626;
}

.retry-button {
  margin-top: 10px;
  padding: 8px 16px;
  background-color: var(--color-background-mute);
  border: 1px solid var(--color-border);
  border-radius: 4px;
  cursor: pointer;
}

.retry-button:hover {
  background-color: var(--color-background-soft);
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(220px, 1fr));
  gap: 20px;
  padding: 20px 0;
}

@media (max-width: 1024px) {
  .items-grid {
    grid-template-columns: repeat(3, minmax(220px, 1fr));
  }
}

@media (max-width: 768px) {
  .items-grid {
    grid-template-columns: repeat(2, minmax(220px, 1fr));
  }
}
@media (max-width: 500px) {
  .items-grid {
    grid-template-columns: 1fr;
  }
}

.item-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
  transition: transform 0.2s ease-in-out;
  display: flex;
  flex-direction: column;
}

.item-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.item-content {
  padding: 15px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.item-title {
  margin: 0 0 10px 0;
  font-size: 1.2rem;
  color: var(--color-heading);
  border-bottom: 1px solid var(--color-border);
  padding-bottom: 8px;
}

.item-details p {
  margin: 8px 0;
  font-size: 0.9rem;
  color: var(--color-text);
  line-height: 1.4;
}

.updated-at {
  font-size: 0.8rem;
  color: var(--color-text-light);
  margin-top: auto;
  font-style: italic;
}
.item-actions {
  margin-top: 10px;
  padding: 10px;
  display: flex;
  justify-content: flex-end;
  gap: 5px;
}

.delete-button {
  padding: 8px 12px;
  background-color: #dc2626;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.3s ease;
}

.delete-button:hover {
  background-color: #b91c1c;
}
.edit-button {
  padding: 8px 12px;
  background-color: #4ade80;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.3s ease;
}

.edit-button:hover {
  background-color: #22c55e;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: var(--color-background-soft);
  padding: 20px;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

.modal h2 {
  margin-bottom: 20px;
  color: var(--color-heading);
}

.modal-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 20px;
}

.modal-content label {
  font-weight: bold;
  color: var(--color-text);
}

.modal-content input {
  padding: 10px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
}
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.modal-actions button {
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
}
.save-button {
  background-color: #4ade80;
  color: white;
  transition: background-color 0.3s ease;
}
.save-button:hover {
  background-color: #22c55e;
}
.cancel-button {
  background-color: #5e5e5f;
  color: var(--color-text);
  transition: background-color 0.3s ease;
}
.cancel-button:hover {
  background-color: #242424;
}

.add-item-container {
  text-align: right;
  margin-bottom: 20px;
}

.add-item-button {
  padding: 10px 15px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.add-item-button:hover {
  background-color: #1d4ed8;
}
</style>
