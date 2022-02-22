<template>
  <h1>{{ $t('navigation.gallery') }}</h1>
  <section>
    <div class="event" v-for="event in imagesByEvent">
      <h2 class="event--title">
        {{ event.title }}
      </h2>
      <ul class="imageContainer">
        <li v-for="(image, index) in event.images">
          <router-link :to="/event/ + event.id + '?img=' + index">
            <img :src="image"></router-link>
        </li>
      </ul>
    </div>
  </section>
</template>

<script setup lang="ts">
let imagesByEvent = [
  {
    title: "Krimidinner",
    id: 1,
    images: [
      "https://unsplash.it/400?random=1",
      "https://unsplash.it/400?random=1",
      "https://unsplash.it/400?random=1",
      "https://unsplash.it/400?random=1",
      "https://unsplash.it/400?random=1"
    ]
  },
  {
    title: "Heimische Kräuterküche!",
    id: 2,
    images: [
      "https://unsplash.it/400",
      "https://unsplash.it/400",
      "https://unsplash.it/400",
      "https://unsplash.it/400",
      "https://unsplash.it/400",
      "https://unsplash.it/400"
    ]
  }
]
</script>

<style scoped lang="scss">
.imageContainer {
  list-style-type: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 4px;

  li {
    aspect-ratio: 1;
    width: 200px;
    height: 200px;
    background-color: white;
    transition: background-color 0.1s ease;

    &:hover {
      background-color: rgba(var(--theme-color-rgb) / 0.7);
    }
    img {
      aspect-ratio: 1;
      width: 200px;
      object-fit: cover;
      mix-blend-mode: multiply;
    }
  }
}
</style>
