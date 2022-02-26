<template>
  <section id="createEvent">
    <h4>{{ $t('eventEditor.title.new') }}</h4>

    <form id="eventForm">
      <input id="motto" name="theme" type="text" placeholder="Motto" required v-model="form.theme">
      <input id="date" name="date" type="datetime-local" required v-model="form.startingDate">
      <input class="menu" name="starter" type="text" placeholder="Vorspeiße" required v-model="form.starter">
      <input class="menu" name="main-dish" type="text" placeholder="Hauptgericht" required v-model="form.mainDish">
      <input class="menu" name="dessert" type="text" placeholder="Nachtisch" required v-model="form.dessert">
      <ImageUpload id="image" v-model="form.imageID"/>
      <textarea name="info" placeholder="Beschreibung" rows="6" required v-model="form.description"></textarea>
      <label id="closing">Anmeldeschluss:<input name="closingDate" type="datetime-local" required v-model="form.closingDate"></label>
    </form>

    <button class="buttonRight" type="submit" @click.prevent="onFormSubmit">Speichern</button>
  </section>
</template>

<script setup lang="ts">
import ImageUpload from "../components/ImageUpload.vue";
import {reactive} from "vue";
import {client} from "../api/client";

const form = reactive({
  theme: '',
  imageID: '',
  startingDate: '',
  closingDate: '',
  starter: '',
  mainDish: '',
  dessert: '',
  description: '',
})

function clearForm() {
  form.theme = ''
  form.imageID = ''
  form.startingDate = ''
  form.closingDate = ''
  form.starter = ''
  form.mainDish = ''
  form.dessert = ''
  form.description = ''
}

async function onFormSubmit() {
  const resp = await client.createEvent({
    'theme': form.theme,
    'imageID': form.imageID,
    'startingDate': new Date(form.startingDate).toISOString(),
    'closingDate': new Date(form.closingDate).toISOString(),
    'starter': form.starter,
    'mainDish': form.mainDish,
    'dessert': form.dessert,
    'description': form.description,
  })

  clearForm()
  console.log(resp.data.id)
}
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
