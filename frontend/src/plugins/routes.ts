import HomeView from '../views/HomeView.vue'
import EventsView from '../views/EventsView.vue'
import EventEditorView from '../views/EventEditorView.vue'
import GalleryView from '../views/GalleryView.vue'
import LegalView from '../views/LegalView.vue'
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
  {path: '/', component: HomeView},
  {path: '/events', component: EventsView},
  {path: '/event/edit', component: EventEditorView},
  {path: '/gallery', component: GalleryView},
  {path: '/legal', component: LegalView},
]

export const router = createRouter({
  history: createWebHistory('/'),
  routes: routes,
})