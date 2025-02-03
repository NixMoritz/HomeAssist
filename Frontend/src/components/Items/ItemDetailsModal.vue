<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>Item Details</h2>
        <button @click="handleClose" class="button-base button-neutral">Close</button>
      </div>
      <div class="modal-content">
        <div class="details-grid">
          <div class="detail-item"><strong>Name:</strong> {{ item?.item_name }}</div>
          <div class="detail-item">
            <strong>Price:</strong> {{ item ? formatPrice(item.item_price) : '' }}
          </div>
          <div class="detail-item">
            <strong>Price per Unit:</strong>
            {{ item ? formatPrice(item.price_per_unit) : '' }}
          </div>
          <div class="detail-item"><strong>Units:</strong> {{ item?.units }}</div>
          <div class="detail-item"><strong>Store Branch:</strong> {{ item?.store_branch }}</div>
          <div class="detail-item"><strong>Weight:</strong> {{ item?.weight }}kg</div>
          <div class="detail-item">
            <strong>Category:</strong> {{ getNullableString(item?.category) }}
          </div>
          <div class="detail-item">
            <strong>Subcategory:</strong> {{ getNullableString(item?.subcategory) }}
          </div>
          <div class="detail-item">
            <strong>Organic:</strong> {{ item?.is_organic ? 'Yes' : 'No' }}
          </div>
          <div class="detail-item">
            <strong>Brand:</strong> {{ getNullableString(item?.brand_name) }}
          </div>
          <div class="detail-item">
            <strong>Barcode:</strong> {{ getNullableString(item?.barcode) }}
          </div>
          <div class="detail-item">
            <strong>Last Updated:</strong>
            {{ item ? formatDate(item.updated_at) : '' }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onBeforeUnmount } from 'vue'
import type { Item } from '@/types/Item'

const props = defineProps<{
  item: Item
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

// Handle cleanup before component unmounts
const handleClose = () => {
  emit('close')
}

onBeforeUnmount(() => {
  // Cleanup any listeners or connections here
})

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
  }).format(price)
}

const formatDate = (date: string): string => {
  return new Date(date).toLocaleString()
}

const getNullableString = (value: string | null | undefined): string => {
  return value ?? '-'
}
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

.modal-header h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--color-heading);
}

.modal-content {
  padding: 24px;
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  padding: 20px;
  width: 100%;
}

.detail-item {
  padding: 12px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

.detail-item strong {
  color: var(--color-heading);
  margin-right: 8px;
}

@media (max-width: 768px) {
  .modal {
    width: 90%;
  }

  .details-grid {
    grid-template-columns: 1fr;
  }

  .modal-header {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
}
</style>
