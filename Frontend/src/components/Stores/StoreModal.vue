<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>{{ isEdit ? 'Edit Store' : 'Add Store' }}</h2>
        <div class="header-actions">
          <button @click="$emit('save')" class="button-base button-primary">
            {{ isEdit ? 'Save Changes' : 'Add Store' }}
          </button>
          <button @click="$emit('close')" class="button-base button-neutral">Cancel</button>
        </div>
      </div>
      <div class="modal-content">
        <div class="form-group">
          <label for="storeName">Store Name:</label>
          <input type="text" id="storeName" v-model="store.store_name" />
        </div>
        <div class="form-group">
          <label for="storeBranch">Store Branch:</label>
          <input type="text" id="storeBranch" v-model="store.store_branch" />
        </div>
        <div class="form-group">
          <label for="storeAddress">Store Address:</label>
          <input type="text" id="storeAddress" v-model="store.store_address" />
        </div>
        <div class="form-group">
          <label for="storePhone">Store Phone:</label>
          <input type="text" id="storePhone" v-model="store.store_phone" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import '@/components/styles/ButtonStyles.vue'
import type { Store } from '@/types/Store.ts'

defineProps<{
  store: Store
  isEdit: boolean
}>()

defineEmits<{
  (e: 'save'): void
  (e: 'close'): void
}>()
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
}
.modal {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  width: 60%;
  max-width: 500px;
  padding: 1rem;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}
.header-actions button {
  margin-left: 0.5rem;
}
.form-group {
  margin-bottom: 1rem;
}
label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
}
input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}
</style>
