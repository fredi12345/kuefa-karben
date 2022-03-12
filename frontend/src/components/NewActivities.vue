<template>
  <section id="activities">
    <h2>Neueste Aktivitäten</h2>
    <div id="flexContainer" >
      <div id="comments">
        <h2>Kommentare</h2>
        <CommentList :comments="comments"/>
      </div>
      <div id="participants">
        <h2>Anmeldungen</h2>
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>

import CommentList from "../views/CommentList.vue";

let comments = [{
  name: "Frederik",
  date: "gestern",
  text: "Ich habe Hunger"
},{
  name: "Frederik",
  date: "gestern",
  text: "Ich habe Hunger und Durst"
}]
</script>

<style scoped lang="scss">
#flexContainer{
  display: flex;
  justify-content: center;
  flex-direction: row;
  gap: 20px;

  div{
    padding: 20px;
    background: var(--card-color);
  }
}
</style>
