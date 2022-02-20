<template>
  <section id="findUs">
    <div id="address">
      <img class="kflogo" src="../assets/logo_wide_thin.png">
      <address>
        Jugendkulturzentrum Karben<br>
        Brunnenstr. 2<br>
        61184 Karben<br>
      </address>
      <a @click="showMap = true" v-if="!showMap">
        <img src="../assets/google-maps.png" alt="Google Maps Logo">
        Google Maps ⧉
      </a>
    </div>
    <div id="socialMedia">
      <h4>Social Media</h4>
      <a href="https://www.facebook.com/kuefakarben/  " target="_blank" rel="noopener"><img
          src="../assets/flogo_RGB_HEX-72.png" alt="Facebook Logo">Facebook</a>
      <a href="https://www.instagram.com/kuefa.karben/" target="_blank" rel="noopener"><img
          src="../assets/insta.svg.png" alt="Instagram Logo">Instagram</a>
    </div>
    <div id="otherSites">
      <h4>Weitere Seiten</h4>
      <a href="https://www.jukuz-karben.de/" target="_blank" rel="noopener">Jugendkulturzentrum Karben</a>
      <a href="https://www.karben.de/" target="_blank" rel="noopener">Stadt Karben</a>
    </div>

  </section>
  <iframe v-if="showMap" id="gmap" src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2551.875012188656!2d8.753086115723688!3d50.23823987944658!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47bd045975035be5%3A0xb426c757ea8e8e98!2sJUKUZ%20Karben%20Jugendkulturzentrum%20Selzerbrunnenhof!5e0!3m2!1sde!2sde!4v1645366663457!5m2!1sde!2sde" width="600" height="450" style="border:0;" loading="lazy"></iframe>
</template>

<script setup lang="ts">
import {ref} from "vue";

const showMap = ref(false);

</script>

<style scoped lang="scss">
#gmap{
  width: 100%;
  height: 60vh;
}

#findUs{
  padding: 20px;
  background-color: var(--theme-color);
  color: white;

  gap: 20px;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  flex-wrap: wrap;


  a{
    font-weight: bold;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    background: white;
    color: var(--theme-color);
    text-decoration: none;
    padding: 6px;
    padding-inline: 10px;
    gap: 10px;
    cursor: pointer;
    margin-block: 10px;

    width: fit-content;

    img{
      max-width: 30px;
      aspect-ratio: 1 / 1;
    }
  }

  >div{
    width: 250px;
    line-height: 1.5;
    text-align: start;
  }

  #address{
    display: flex;
    flex-direction: column;
    line-height: 2;
    gap: 8px;
    address{
      font-style: normal;
    }
    .kflogo{
      margin-top: 20px;
      width: 250px;
      height: 50px;
      object-fit: cover;
    }
  }
}
</style>
