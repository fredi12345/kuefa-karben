<template>
  <section>
    <h2>Aktuelle Veranstaltungen</h2>
    <div id="events">
      <EventCard v-for="event in events" :event="event"></EventCard>
      <router-link to="/events" id="showAll">></router-link>
    </div>
  </section>
</template>
<script setup lang="ts">
import EventCard from "./EventCard.vue";

let events = [
  {
    title: "Heimische Kräuterküche",
    id: 1,
    date: "22.02.2022",
    img: "https://unsplash.it/640/425",
    description: "Lord Grey bittet zum Gärtnerkongress! Seinem Ruf folgen Botaniker, Lifestyle-Blogger und Imker, um sich über die neusten Erfindungen der Branche auszutauschen. Doch schon bald müssen sich die Gäste einer jahrzehntealten Frage stellen: Ist der Mörder immer der Gärtner?"
  },
  {
    title: "Krimidinner - To Bee or Not To Bee or Not To Bee or Not To Bee!",
    date: "25.02.2022",
    id: 2,
    img: "https://unsplash.it/640/425?random=1",
    description: "Lord Grey bittet zum Gärtnerkongress! Seinem Ruf folgen Botaniker, Lifestyle-Blogger und Imker, um sich über die neusten Erfindungen der Branche auszutauschen. Doch schon bald müssen sich die Gäste einer jahrzehntealten Frage stellen: Ist der Mörder immer der Gärtner?"
  }
]
</script>
<style scoped lang="scss">
section{
  margin-block: 20px;
}
#events {
  display: flex;
  flex-direction: row;
  gap: 10px;
  justify-content: center;
}

#showAll{
  font-size: 3rem;
  text-decoration: none;
  border-radius: 50%;
  color: var(--theme-color);
  border: 3px solid var(--theme-color);
  width: 100px;
  aspect-ratio: 1;
  flex-shrink: 0;
  align-self: center;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: 20px;
}

</style>
