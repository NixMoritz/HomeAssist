<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>Receipt Details</h2>
        <button @click="$emit('close')" class="button-base button-neutral">Close</button>
      </div>
      <div class="modal-content">
        <div class="details-grid">
          <div class="detail-item"><strong>Receipt #:</strong> {{ receipt?.receipt_id }}</div>
          <div class="detail-item">
            <strong>Date Issued:</strong> {{ formatDate(receipt?.date_issued) }}
          </div>
          <div class="detail-item">
            <strong>Store:</strong>
            {{
              receipt?.store_id
                ? receipt.store_name + ' (' + receipt.store_branch + ')'
                : 'No Store Selected'
            }}
          </div>
          <div class="detail-item">
            <strong>Address:</strong>
            {{ receipt?.store_id ? receipt.store_address : '-' }}
          </div>
          <div class="detail-item">
            <strong>Total Amount:</strong> {{ formatPrice(receipt?.total_amount) }}
          </div>
          <div class="detail-item">
            <strong>Discount:</strong> {{ formatPrice(receipt?.total_discount_amount) }}
          </div>
          <div class="detail-item">
            <strong>Net Amount:</strong> {{ formatPrice(receipt?.net_amount) }}
          </div>
          <div class="detail-item">
            <strong>Tax Amount:</strong> {{ formatPrice(receipt?.tax_amount) }}
          </div>
          <div class="detail-item">
            <strong>Payment Method:</strong> {{ receipt?.payment_method }}
          </div>
          <div class="detail-item"><strong>Type:</strong> {{ receipt?.receipt_type }}</div>
          <div class="detail-item">
            <strong>Loyalty Card:</strong> {{ receipt?.loyalty_card_number || '-' }}
          </div>
          <div class="detail-item">
            <strong>Cashier:</strong> {{ receipt?.cashier_name || '-' }}
          </div>
          <div class="detail-item">
            <strong>Receipt Number:</strong> {{ receipt?.receipt_number || '-' }}
          </div>
          <div class="detail-item"><strong>Notes:</strong> {{ receipt?.notes || '-' }}</div>
          <div class="detail-item">
            <strong>Last Updated:</strong> {{ formatDate(receipt?.updated_at) }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import type { Receipt } from '@/types/Receipt'

const props = defineProps<{
  receipt: Receipt | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const formatDate = (date?: string): string => {
  return date ? new Date(date).toLocaleString() : '-'
}

const formatPrice = (price?: number): string => {
  if (price === undefined || price === null) return '-'
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(price)
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
  max-height: 90vh; /* Ensure modal does not exceed viewport height */
  overflow-y: auto; /* Enable scrolling */
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
  .modal-header {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
  .details-grid {
    grid-template-columns: 1fr;
  }
}
</style>
