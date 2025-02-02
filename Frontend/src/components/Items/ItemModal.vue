<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>{{ isEdit ? 'Edit Item' : 'Add New Item' }}</h2>
        <div class="header-actions">
          <button @click="$emit('save')" class="button-base button-primary">
            {{ isEdit ? 'Save Changes' : 'Add Item' }}
          </button>
          <button @click="$emit('close')" class="button-base button-neutral">Cancel</button>
        </div>
      </div>

      <div class="modal-content">
        <div class="info-card">
          <div class="form-grid">
            <div class="form-group">
              <label for="itemName">Name:</label>
              <input type="text" id="itemName" v-model="item.item_name" />
            </div>

            <div class="form-group">
              <label for="unitPrice">Price:</label>
              <input type="number" id="unitPrice" v-model.number="item.unit_price" step="0.01" />
            </div>

            <div class="form-group">
              <label for="units">Units:</label>
              <input type="number" id="units" v-model.number="item.units" />
            </div>

            <div class="form-group">
              <label for="weight">Weight (kg):</label>
              <input type="number" id="weight" v-model.number="item.weight" step="0.01" />
            </div>

            <div class="form-group">
              <label for="category">Category:</label>
              <input type="text" id="category" v-model="item.category" />
            </div>

            <div class="form-group">
              <label for="subcategory">Subcategory:</label>
              <input type="text" id="subcategory" v-model="item.subcategory" />
            </div>

            <div class="form-group">
              <label for="brandName">Brand:</label>
              <input type="text" id="brandName" v-model="item.brand_name" />
            </div>

            <div class="form-group">
              <label for="storeBranch">Store Branch:</label>
              <input type="text" id="storeBranch" v-model="item.store_branch" />
            </div>

            <div class="form-group">
              <label for="isOrganic">Organic:</label>
              <select id="isOrganic" v-model="item.is_organic">
                <option :value="null">Unknown</option>
                <option :value="true">Yes</option>
                <option :value="false">No</option>
              </select>
            </div>

            <div class="form-group">
              <label for="barcode">Barcode:</label>
              <input type="text" id="barcode" v-model="item.barcode" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import '@/components/styles/ButtonStyles.vue'
import type { Item } from '@/types/Item'

defineProps<{
  item: Item
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
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  width: 60%;
  max-width: 800px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  position: sticky;
  top: 0;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 1;
}

.modal-content {
  padding: 24px;
}

.info-card {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 20px;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: var(--color-heading);
  font-size: 0.9rem;
}

.form-group input,
.form-group select {
  padding: 10px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-text);
  font-size: 0.95rem;
  transition: all 0.2s ease;
}

.form-group input:focus,
.form-group select:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.4);
  background: rgba(255, 255, 255, 0.1);
}

@media (max-width: 768px) {
  .modal {
    width: 90%;
  }

  .form-grid {
    grid-template-columns: 1fr;
  }

  .modal-header {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
}
</style>
