<template>
  <div>
    <h1>{{ event.theme }}</h1>
    <p>
      <time >{{ event.startingDate }}</time>
    </p>
    <img id="mainImg" :src="event.img"/>
    <div id="menu">
      <p class="menuItem menu__1">{{ event.starter }}</p>
      <p class="menuItem menu__2">{{ event.mainDish }}</p>
      <p class="menuItem menu__3">{{ event.dessert }}</p>
    </div>
    <p>{{ event.description }}</p>
    <div id="images" v-if="event.images && event.images.length > 0">
      <h2>Bilder</h2>
      <ul class="imageContainer">
        <li v-for="(image, index) in event.images">
          <img :src="image" @click="this.$refs.viewer.viewImage(index)">
        </li>
      </ul>
    </div>
    <ImageViewer ref="viewer" :images="event.images"/>
  </div>
</template>
<script setup lang="ts">

import {CreateEventRequest} from "../api/generated";
import ImageViewer from "./ImageViewer.vue";
import {ref} from "vue";

defineProps<{
  event: CreateEventRequest
}>()
const viewer = ref(null)
function doDas(){
  console.log(this.$refs)
}
</script>
<style scoped lang="scss">

#mainImg {
  max-width: 100%;
}

#menu {
  .menuItem {
    padding-inline: 40px;
  }

  .menu__1 {
    text-align: start;
  }

  .menu__2 {
    text-align: center;
  }

  .menu__3 {
    text-align: end;
  }
}

.imageContainer {
  list-style-type: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 10px;
  padding: 0;

  li {
    display: block;
    position: relative;
    width: 150px;
    aspect-ratio: 1;
    background-color: white;
    transition: background-color 0.1s ease;

    &:hover {
      background-color: rgba(var(--theme-color-rgb) / 0.7);
    }

    img {
      display: block;
      height: 100%;
      width: 100%;
      object-fit: cover;
      mix-blend-mode: multiply;
    }
  }
}
</style>
