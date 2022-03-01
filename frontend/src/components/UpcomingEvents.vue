<template>
  <section>
    <h2>Aktuelle Veranstaltungen</h2>
    <div id="events">
      <EventCard
          v-for="event in events.value"
          :id="event.id"
          :title="event.theme"
          :image-url="event.thumbnailURL"
          :date="event.date"
      ></EventCard>
      <router-link to="/events" id="showAll">></router-link>
    </div>
  </section>
</template>
<script setup lang="ts">
import EventCard from "./EventCard.vue";
import {client} from "../api/client";
import {onMounted, reactive} from "vue";
import {EventTeaser} from "../api/generated";

let events = reactive({value: Array<EventTeaser>()});
onMounted(async () => {
  const resp = await client.getEvents(2)
  // console.log(resp.data);
  // for (const event of resp.data.events) {
  //   events.push(event)
  // }
  events.value = resp.data.events
})

</script>
<style scoped lang="scss">
section {
  margin-block: 20px;
}

#events {
  display: flex;
  flex-direction: row;
  gap: 10px;
  justify-content: center;
}

#showAll {
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
