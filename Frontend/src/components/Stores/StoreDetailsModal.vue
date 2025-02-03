<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>Store Details</h2>
        <button @click="$emit('close')" class="button-base button-neutral">Close</button>
      </div>
      <div class="modal-content">
        <div class="details-grid">
          <div class="detail-item"><strong>Name:</strong> {{ store?.store_name }}</div>
          <div class="detail-item"><strong>Branch:</strong> {{ store?.store_branch }}</div>
          <div class="detail-item"><strong>Address:</strong> {{ store?.store_address }}</div>
          <div class="detail-item"><strong>Phone:</strong> {{ store?.store_phone }}</div>
          <div class="detail-item"><strong>Type:</strong> {{ store?.store_type || '-' }}</div>
          <div class="detail-item">
            <strong>Category:</strong> {{ store?.store_category || '-' }}
          </div>
          <div class="detail-item">
            <strong>Last Updated:</strong> {{ formatDate(store?.updated_at) }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Store } from '@/types/Store'

const props = defineProps<{
  store: Store | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const formatDate = (date: string | undefined): string => {
  return date ? new Date(date).toLocaleString() : '-'
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
  max-width: 500px;
  padding: 1rem;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.detail-item {
  padding: 12px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

@media (max-width: 768px) {
  .modal {
    width: 90%;
  }
  .details-grid {
    grid-template-columns: 1fr;
  }
}
</style>
