<!-- TopBar.vue -->
<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <nav class="top-bar">
    <div class="container">
      <div class="logo-section">
        <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="40" height="40" />
        <span class="site-name">HomeAssist</span>
      </div>
      <div class="nav-links">
        <RouterLink to="/" class="nav-link">Home</RouterLink>
        <RouterLink to="/about" class="nav-link">About</RouterLink>
        <RouterLink to="/contact" class="nav-link">Contact</RouterLink>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.top-bar {
  background-color: #aa6061;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0.5rem 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logo {
  height: 40px;
  width: 40px;
}

.site-name {
  font-size: 1.55rem;
  font-weight: 600;
  color: #2c3e50;
}

.nav-links {
  display: flex;
  gap: 1.5rem;
}

.nav-link {
  color: #2c3e50;
  text-decoration: none;
  font-size: 1rem;
  padding: 0.5rem;
  transition: color 0.2s;
}

.nav-link:hover {
  color: #ffffff;
}

.router-link-active {
  color: #fff000;
  font-weight: 600;
}

@media (max-width: 768px) {
  .container {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }

  .nav-links {
    width: 100%;
    justify-content: center;
  }
}
</style>
