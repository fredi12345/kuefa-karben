<template>
  <section id="hero" class="hero">
    <div class="text">
      <h1>Küche für alle</h1>
      <p>Gutes Essen in
        netter Runde.</p>
    </div>
  </section>
  <AboutKuefa/>
  <UpcomingEvents/>
  <FindUs/>
</template>

<script setup lang="ts">
import AboutKuefa from "../components/AboutKuefa.vue";
import FindUs from "../components/FindUs.vue";
import UpcomingEvents from "../components/UpcomingEvents.vue";


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
    left: max(0%, calc(25% - 200px));
    padding: 20px 80px;

    &:before {
      z-index: 0;
      content: '';
      mix-blend-mode: multiply;
      position: absolute;
      inset: 0;
      background-color: var(--theme-color);
    }

    p, h1 {
      position: relative;
    }

    h1 {
      font-size: 2rem;
    }
  }
}
</style>
