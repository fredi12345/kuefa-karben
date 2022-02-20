<template>
  <AboutKuefa/>
  <section id="hero" class="hero">
    <div class="text">
      <h1>Küche für alle</h1>
      <p>Gutes Essen in
        netter Runde.</p>
    </div>

  </section>

  <h2>Aktuelle Veranstaltungen</h2>
  <div id="events">
    <EventCard v-for="event in events" :event="event"></EventCard>
  </div>
</template>

<script setup lang="ts">
import EventCard from "../components/EventCard.vue";
import AboutKuefa from "../components/AboutKuefa.vue";

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

let count = 1;
let heros = [
  "https://xn--kfa-karben-9db.de/public/images/1583881459088017073.jpg",
  "https://xn--kfa-karben-9db.de/public/images/1583876349736466949.jpg",
  "https://xn--kfa-karben-9db.de/public/images/1561540425463947945.jpg"
]
setInterval(() => {
  document.getElementById("hero").style.backgroundImage = `url("${heros[count++ % 3]}")`;
}, 6000)
</script>

<style scoped lang="scss">
#events {
  display: flex;
  flex-direction: row;
  gap: 10px;
  justify-content: center;
}

.hero {
  position: relative;
  height: 50vh;
  background-size: cover;
  background-position: center center;
  background-image: url("https://xn--kfa-karben-9db.de/public/images/1583876349736466949.jpg");
  transition: background-image 2s ease;

  .text {
    position: absolute;
    margin: 20px 20px;
    color: white;
    bottom: 0;
    padding: 20px 80px;

    &:before {
      z-index: 0;
      content: '';
      mix-blend-mode: multiply;
      position: absolute;
      inset: 0;
      background-color: var(--theme-color);
    }

    p, h1{
      position: relative;
    }
    h1{
      font-size: 2rem;
    }
  }
}
</style>
