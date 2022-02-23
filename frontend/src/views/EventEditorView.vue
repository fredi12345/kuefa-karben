<template>
  <section id="createEvent">
    <h4>{{ $t('eventEditor.title.new') }}</h4>

    <ImageUpload v-model="imageID"/>


    <form id="eventForm">
      <input id="motto" name="theme" type="text" placeholder="Motto" required>
      <input id="date" name="date" type="datetime-local" required>
      <input class="menu" name="starter" type="text" placeholder="Vorspeiße" required>
      <input class="menu" name="main-dish" type="text" placeholder="Hauptgericht" required>
      <input class="menu" name="dessert" type="text" placeholder="Nachtisch" required>
      <textarea name="info" placeholder="Beschreibung" rows="6" required></textarea>
      <label id="closing">Anmeldeschluss:<input name="closingDate"  type="datetime-local" required></label>
      <input id="image" name="image" type="file" accept="image/*" required>
    </form>
    <button class="buttonRight" type="submit">Speichern</button>
  </section>
</template>

<script setup lang="ts">
import ImageUpload from "../components/ImageUpload.vue";
import {ref} from "vue";

const imageID = ref('')
</script>

<style scoped lang="scss">
section{
  --gap: 6px;
  --inp-wdt: 250px;
  padding-inline: 3%;
}
#eventForm{
  margin-top: 50px;
  position: relative;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: var(--gap);

  input{
    height: 40px;
    box-sizing: border-box;
  }
  input, textarea{
    min-width: var(--inp-wdt);
    flex-grow: 1;
  };

  textarea{
    flex-basis: 100%;
    text-align: center;
  }

  #motto, #date, label, #image {
    flex-basis: 48%;
  }
  label{
    padding-inline-start: 6px;
    display: flex;
    align-items: center;
    white-space: nowrap;
    flex-grow: 1;
    input{
      flex-grow: 1;
      margin-left: 10px;
    }
  }
}
button{
  margin-top: var(--gap);
  display: block;
  margin-left: auto;
}
@media (max-width: 1200px) {
  .menu{
    flex-basis: 100%;
  }
}
</style>
