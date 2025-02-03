<template>
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2>{{ isEdit ? 'Edit Receipt' : 'Add New Receipt' }}</h2>
        <div class="header-actions">
          <button @click="$emit('save')" class="button-base button-primary">
            {{ isEdit ? 'Save Changes' : 'Add Receipt' }}
          </button>
          <button @click="$emit('close')" class="button-base button-neutral">Cancel</button>
        </div>
      </div>

      <div class="modal-content">
        <!-- Store selection card -->
        <div class="info-card store-selection">
          <label>Selected Store:</label>
          <div class="selected-store">
            <span v-if="receipt.store_id">
              {{ receipt.store_name }} - {{ receipt.store_branch }}
            </span>
            <span v-else>No store selected</span>
          </div>
          <button @click="$emit('openStoreSelect')" class="button-base button-special">
            Select Store
          </button>
        </div>

        <!-- Amount details card -->
        <div class="info-card">
          <h3>Amount Details</h3>
          <div class="form-grid">
            <div class="form-group">
              <label for="totalAmount">Total Amount:</label>
              <input
                type="number"
                id="totalAmount"
                v-model.number="receipt.total_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="discountAmount">Discount:</label>
              <input
                type="number"
                id="discountAmount"
                v-model.number="receipt.total_discount_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="netAmount">Net Amount:</label>
              <input type="number" id="netAmount" v-model.number="receipt.net_amount" step="0.01" />
            </div>

            <div class="form-group">
              <label for="taxAmount">Tax Amount:</label>
              <input type="number" id="taxAmount" v-model.number="receipt.tax_amount" step="0.01" />
            </div>
          </div>
        </div>

        <!-- Receipt details card -->
        <div class="info-card">
          <h3>Receipt Details</h3>
          <div class="form-grid">
            <div class="form-group">
              <label for="datetime">Date and Time:</label>
              <input type="datetime-local" id="datetime" v-model="receipt.date_issued" />
            </div>

            <div class="form-group">
              <label for="paymentMethod">Payment Method:</label>
              <select id="paymentMethod" v-model="receipt.payment_method">
                <option v-for="(value, key) in PAYMENT_METHODS" :key="key" :value="value">
                  {{ key }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label for="receiptType">Receipt Type:</label>
              <select id="receiptType" v-model="receipt.receipt_type">
                <option v-for="(value, key) in RECEIPT_TYPES" :key="key" :value="value">
                  {{ key }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label for="receiptNumber">Receipt Number:</label>
              <input type="text" id="receiptNumber" v-model="receipt.receipt_number" />
            </div>
          </div>
        </div>

        <!-- Additional details card -->
        <div class="info-card">
          <h3>Additional Details</h3>
          <div class="form-grid">
            <div class="form-group">
              <label for="loyaltyCard">Loyalty Card:</label>
              <input type="text" id="loyaltyCard" v-model="receipt.loyalty_card_number" />
            </div>

            <div class="form-group">
              <label for="cashierName">Cashier:</label>
              <input type="text" id="cashierName" v-model="receipt.cashier_name" />
            </div>

            <div class="form-group full-width">
              <label for="notes">Notes:</label>
              <textarea id="notes" v-model="receipt.notes" rows="3"></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import type { Receipt } from '@/types/Receipt'
import { PAYMENT_METHODS, RECEIPT_TYPES } from '@/types/ReceiptTypes'
import '@/components/styles/ButtonStyles.vue'

defineProps<{
  receipt: Receipt
  isEdit: boolean
}>()

defineEmits<{
  (e: 'save'): void
  (e: 'close'): void
  (e: 'openStoreSelect'): void
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
  z-index: 1000;
}

.modal {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  width: 60%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  padding: 1rem;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.modal-header h2 {
  margin: 0;
  font-size: 1.5rem;
  color: var(--color-heading);
}

.header-actions {
  display: flex;
  gap: 12px;
}

.modal-content {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 0;
}

.info-card h3 {
  margin: 0 0 16px 0;
  font-size: 1.1rem;
  color: var(--color-heading);
  opacity: 0.9;
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

.form-group.full-width {
  grid-column: span 2;
}

.form-group label {
  color: var(--color-heading);
  font-size: 0.9rem;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 10px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-text);
  font-size: 0.95rem;
  transition: all 0.2s ease;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.4);
  background: rgba(255, 255, 255, 0.1);
}

.store-selection {
  margin-bottom: 20px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

.selected-store {
  margin: 8px 0;
  padding: 8px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
}

.select-store-button {
  background: rgba(59, 130, 246, 0.2);
  color: #60a5fa;
  padding: 8px 16px;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.select-store-button:hover {
  background: rgba(59, 130, 246, 0.3);
}

.save-button {
  background: rgba(59, 130, 246, 0.2);
  color: #60a5fa;
}

.cancel-button {
  background: rgba(107, 114, 128, 0.2);
  color: #9ca3af;
}

.save-button:hover {
  background: rgba(59, 130, 246, 0.3);
}

.cancel-button:hover {
  background: rgba(107, 114, 128, 0.3);
}

/* Add responsive design */
@media (max-width: 1024px) {
  .modal {
    width: 80%;
  }
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
}
</style>
