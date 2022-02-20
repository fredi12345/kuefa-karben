<template>
  <section>
    <h2>Aktuelle Veranstaltungen</h2>
    <div id="events">
      <EventCard v-for="event in events" :event="event"></EventCard>
    </div>
  </section>
</template>
<script setup lang="ts">
import EventCard from "./EventCard.vue";

let events = [
  {
    title: "Heimische Kräuterküche",
    date: "22.02.2022",
    img: "https://unsplash.it/640/425",
    description: "Lord Grey bittet zum Gärtnerkongress! Seinem Ruf folgen Botaniker, Lifestyle-Blogger und Imker, um sich über die neusten Erfindungen der Branche auszutauschen. Doch schon bald müssen sich die Gäste einer jahrzehntealten Frage stellen: Ist der Mörder immer der Gärtner?"
  },
  {
    title: "Krimidinner - To Bee or Not To Bee or Not To Bee or Not To Bee!",
    date: "25.02.2022",
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

</style>
