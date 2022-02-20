<template>
  <header>
    <img alt="Küfa Logo" src="../assets/logo_wide_thin.png">
  </header>
  <nav class="navigation">
    <router-link to="/">{{ $t('navigation.home') }}</router-link>
    <router-link to="/events">{{ $t('navigation.events') }}</router-link>
    <router-link to="/gallery">{{ $t('navigation.gallery') }}</router-link>
    <router-link to="/event/edit">{{ $t('navigation.createEvent') }}</router-link>
  </nav>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
header img {
  width: 100%;
  max-width: 1000px;
  height: 20vh;
  filter: drop-shadow(4px 4px 3px black);
  object-fit: cover;
}

nav {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
  padding: 10px;
  border-radius: 10px;
  background-color: var(--theme-color);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.19), 0 6px 6px rgba(0, 0, 0, 0.23);
  text-shadow: 1px 1px 1px rgba(0 0 0 / 0.5);
}

nav a {
  padding: 8px 16px;
  transition: background-color 0.15s;
  text-decoration: none;
  text-transform: uppercase;
  color: white;
  border-radius: 5px;
}

nav a:hover, nav a.router-link-active {
  transition: background-color 0.30s;
  background-color: rgba(255, 255, 255, 0.2);
}
</style>
