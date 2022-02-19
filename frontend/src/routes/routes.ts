import HomeView from '../views/HomeView.vue'
import EventsView from '../views/EventsView.vue'
import GalleryView from '../views/GalleryView.vue'
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
  {path: '/', component: HomeView},
  {path: '/events', component: EventsView},
  {path: '/gallery', component: GalleryView},
]

const router = createRouter({
  history: createWebHistory('/'),
  routes: routes,
})

export default router