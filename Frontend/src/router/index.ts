import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ItemsView from '@/views/ItemsView.vue'
import DashboardView from '@/views/DashboardView.vue'
import ReceiptsView from '@/views/ReceiptsView.vue'
import StoreGrid from '@/components/Stores/StoreGrid.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/items',
      name: 'items',
      component: ItemsView,
    },
    {
      path: '/stores',
      name: 'stores',
      component: StoreGrid,
    },
    {
      path: '/receipts',
      name: 'receipts',
      component: ReceiptsView,
    },
  ],
})

export default router
