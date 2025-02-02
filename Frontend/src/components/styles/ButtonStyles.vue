<template>
  <!-- This is a style-only component -->
  <div style="display: none"></div>
</template>

<script setup lang="ts">
// No script needed
</script>

<style>
.button-base {
  background-color: #ffffff;
  border: 1px solid #222222;
  border-radius: 8px;
  box-sizing: border-box;
  color: #222222;
  cursor: pointer;
  display: inline-block;
  font-family:
    Circular,
    -apple-system,
    BlinkMacSystemFont,
    Roboto,
    'Helvetica Neue',
    sans-serif;
  font-size: 16px;
  font-weight: 600;
  line-height: 20px;
  margin: 0;
  outline: none;
  padding: 13px 23px;
  position: relative;
  text-align: center;
  text-decoration: none;
  touch-action: manipulation;
  transition:
    box-shadow 0.2s,
    transform 0.1s;
  user-select: none;
  width: auto;
}

.button-base:focus-visible {
  box-shadow:
    #222222 0 0 0 2px,
    rgba(255, 255, 255, 0.8) 0 0 0 4px;
  transition: box-shadow 0.2s;
}

.button-base:active {
  background-color: #f7f7f7;
  border-color: #000000;
  transform: scale(0.96);
}

.button-base:disabled {
  border-color: #dddddd;
  color: #dddddd;
  cursor: not-allowed;
  opacity: 1;
}

/* Primary action buttons */
.button-base.button-primary {
  background-color: #4caf50;
  border-color: #45a049;
  color: white;
}

/* Secondary action buttons */
.button-base.button-secondary {
  background-color: #2196f3;
  border-color: #1e88e5;
  color: white;
}

/* Danger action buttons */
.button-base.button-danger {
  background-color: #f44336;
  border-color: #e53935;
  color: white;
}

/* Special action buttons */
.button-base.button-special {
  background-color: #9c27b0;
  border-color: #8e24aa;
  color: white;
}

/* Cancel/neutral buttons */
.button-base.button-neutral {
  background-color: #6b7280;
  border-color: #4b5563;
  color: white;
}

/* Hover states */
.button-base.button-primary:hover,
.button-base.button-secondary:hover,
.button-base.button-danger:hover,
.button-base.button-special:hover,
.button-base.button-neutral:hover {
  opacity: 0.9;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .button-base {
    padding: 10px 18px;
    font-size: 14px;
  }
}
</style>
