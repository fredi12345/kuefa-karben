<template>
  <router-link :to="`/event/${id}`" class="event">
    <img :src="imageUrl" alt="Veranstaltungsbild">
    <div class="details">
      <h4>{{ title }}</h4>
      <time :datetime="date">{{ date }}</time>
    </div>
  </router-link>
</template>

<script setup>
defineProps({
  id: String,
  title: String,
  imageUrl: String,
  date: String,
})
</script>

<style scoped lang="scss">
.event {
  display: flex;
  flex-direction: column;
  max-width: 400px;
  text-align: start;

  text-decoration: none;
  color: currentColor;

  background-color: white; /* yes white, not card-color! is (visually) overwritten in .details */
  box-shadow: 0 0 10px 0 rgb(0 0 0 / 15%);

  transition-property: box-shadow, background-color;
  transition-duration: 0.3s;

  &:hover {
    background-color: rgba(var(--theme-color-rgb) / 0.7);
    box-shadow: 0 0 30px 0 rgb(0 0 0 / 25%);
  }

  img {
    mix-blend-mode: multiply;
    width: 100%;
    aspect-ratio: 2 / 1;
    object-fit: cover;
  }

  .details {
    border-top: 4px solid var(--theme-color);
    flex-grow: 1;
    background-color: var(--card-color);
    padding: 10px 16px 16px;

    h4 {
      margin: 0;
    }

    time {
      display: block;
      margin-bottom: 16px;
    }

    p {
      display: -webkit-box;
      margin: 0;
      -webkit-line-clamp: 3;
      -webkit-box-orient: vertical;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
}
</style>
