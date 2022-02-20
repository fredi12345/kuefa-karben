<template>
  <section id="createEvent">
    <h4>{{ $t('eventEditor.title.new') }}</h4>

    <ImageUpload v-model="imageID"/>


    <!--    <form>-->
    <!--      <input name="theme" type="text" placeholder="Motto" required>-->
    <!--      <input name="date" type="date" required>-->
    <!--      <input name="starter" type="text" placeholder="Vorspeiße" required>-->
    <!--      <input name="main-dish" type="text" placeholder="Hauptgericht" required>-->
    <!--      <input name="dessert" type="text" placeholder="Nachtisch" required>-->
    <!--      <input name="image" type="file" accept="image/*" required>-->
    <!--      <textarea name="info" placeholder="Beschreibung" rows="6" required></textarea>-->
    <!--      <label>Anmeldeschluss:<input name="closingDate" required></label>-->
    <!--      <input class="buttonRight" type="submit" value="Speichern">-->
    <!--    </form>-->
  </section>
</template>

<script setup lang="ts">
import ImageUpload from "../components/ImageUpload.vue";
import {ref} from "vue";

const imageID = ref('')
</script>

<style scoped lang="scss">

</style>