<template>
  <section class="event">
    <main>
      <h1>{{ event.title }}</h1>
      <img :src="event.img"/>

      <div id="menu">
        <p class="menuItem menu__1">{{ event.vorspeise }}</p>
        <p class="menuItem menu__2">{{ event.hauptgang }}</p>
        <p class="menuItem menu__3">{{ event.nachtisch }}</p>
      </div>
      <p>{{ event.description }}</p>
      <div id="images">
        <h2>Bilder</h2>
        <ul class="imageContainer">
          <li v-for="(image, index) in event.images">
            <img :src="image" @click="viewImage(index)">
          </li>
        </ul>
      </div>
    </main>
    <aside>
      <div id="registration">
        <h2>Anmeldung</h2>
        <p class="info closing">Anmeldeschluss: {{ event.closingDate }}</p>
        <form>
          <input class="name" type="text" placeholder="Name" required><br>
          <textarea rows="3" name="message" id="participationMessage" type="text"
                    placeholder="Bemerkung/Nachricht (Optional)"
                    maxlength="1000"></textarea>
          <label class="menuCount"><input type="number" name="classic_count" value="1" min="0">x Klassisch</label>
          <label class="menuCount"><input type="number" name="vegetarian_count" value="0" min="0">x Vegetarisch</label>
          <label class="menuCount"><input type="number" name="vegan_count" value="0" min="0">x Vegan</label>
          <p class="info">Dein Name erscheint öffentlich auf dieser Seite. Benutze gegebenenfalls einen
            Spitznamen, um anonym zu bleiben.</p>
          <button type="submit" class="">Teilnehmen</button>
        </form>
      </div>
      <div id="comments">
        <h2>Kommentare</h2>
        <ol id="commentList">
          <li class="comment" v-for="comment in event.comments">
            <span class="comment--author">{{ comment.name }} </span>
            <span class="comment--date"> {{ comment.date }}</span>
            <p class="comment--text">{{ comment.text }}</p>
          </li>
        </ol>
        <form id="commentForm">
          <textarea rows="3" placeholder="Kommentar"></textarea>
          <button>Senden</button>
        </form>
      </div>
    </aside>
    <div id="imageViewer" v-if="viewerVisible" @click="closeViewer">
      <div class="actions" @click.stop>
        <button @click="closeViewer">Close</button>
        <button @click="changeImage(-1)">Previous</button>
        <button @click="changeImage(1)">Next</button>
      </div>
      <img :src="event.images[bigImgIndex]" @click.stop/>
    </div>
  </section>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useRoute} from "vue-router";
import {router} from "../plugins/routes";

const event = {
  title: "Heimische Kräuterküche",
  date: "22.02.2022",
  img: "https://unsplash.it/640/425",
  description: "Lord Grey bittet zum Gärtnerkongress! Seinem Ruf folgen Botaniker, Lifestyle-Blogger und Imker, um sich über die neusten Erfindungen der Branche auszutauschen. Doch schon bald müssen sich die Gäste einer jahrzehntealten Frage stellen: Ist der Mörder immer der Gärtner?",
  vorspeise: "Handkäs-Bärlauch-Dip mit geröstetem Brot",
  hauptgang: "Grüne Soße mit Kartoffeln und Ei",
  nachtisch: "Zitronen-Basilikum-Sorbet",
  closingDate: "21.02.2022",
  comments: [{
    name: "Fredi",
    date: "22.02.2022",
    text: "Aufgrund der Corona-Krise ist die Veranstaltung auf unbestimmte Zeit verschoben und es sind deshalb derzeit auch keine Anmeldungen möglich."
  },
    {
      name: "Si.",
      date: "10.03.2019",
      text: "Ich möchte auch gern zum Krimidinner kommen. Wann gibt es so etwas nochmal?"
    }

  ],
  images: [
    "https://unsplash.it/400?random=0",
    "https://unsplash.it/500?random=1",
    "https://unsplash.it/400/200?random=2",
    "https://unsplash.it/400/800?random=3",
    "https://unsplash.it/400?random=4"
  ]
}

const viewerVisible = ref(false);
const bigImgIndex = ref(0);

const route = useRoute();
if (route.query.img && route.query.img !== "") {
  viewImage(Number(route.query.img));
}

function viewImage(index: number) {
  bigImgIndex.value = index;
  viewerVisible.value = true;
}

function changeImage(to: number) {
  //previous to=-1
  //next to=1
  bigImgIndex.value += to;
  if (bigImgIndex.value < 0) bigImgIndex.value = event.images.length - 1;
  if (bigImgIndex.value >= event.images.length) bigImgIndex.value = 0;
  //TODO darf man so machen?
  router.replace({path: route.path, query: {img: bigImgIndex.value.toString()}})
}

function closeViewer() {
  viewerVisible.value = false;
  router.replace(route.path);
}
</script>

<style scoped lang="scss">
.event {
  display: flex;
  flex-direction: row;
  gap: 40px;
  padding-inline: 3%;
  flex-wrap: wrap;
}

main, aside{
  width: 400px;
}
main{
  flex: 4;
  flex-basis: 500px;
}
aside{
  flex: 1;
  flex-basis: 350px;
}
main img{
  max-width: 100%;
}
aside {
  --padding: 20px;
  div {
    padding-bottom: var(--padding);
    overflow: auto;
    background-color: var(--card-color);
    text-align: start;
    margin-top: 20px;
  }
  input:not(.menuItem), textarea {
    margin-top: 6px;
    box-sizing: border-box;
    width: 100%;
  }


  button{
    float: right;
  }
}

#menu {
  .menuItem {

    padding-inline: 40px;
  }

  .menu__1 {
    text-align: start;
  }

  .menu__2 {
    text-align: center;
  }

  .menu__3 {
    text-align: end;
  }
}

#registration {
  padding-inline: var(--padding);
  .menuCount {
    display: block;
    padding-inline-start: 20px;

    input {
      width: 50px;
      margin-inline-end: 6px;
    }
  }
}

#comments {
  padding-inline: var(--padding);
  #commentList {
    list-style: none;
    padding: 0;
  };

  .comment{
    overflow: auto;
    margin-inline: calc(-1 * var(--padding));
    padding-inline: calc(var(--padding));
    padding-top: var(--padding);
    &:nth-child(even){
      background-color: rgba(var(--theme-color-rgb) / 0.2);
    }

    &--author, &--date{
      opacity: 0.7;
    }
    &--date{
      float: right;
    }
  }

  #commentForm{
    textarea {
      box-sizing: border-box;
      width: 100%;
    }
  }
}

.imageContainer {
  list-style-type: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 10px;
  padding: 0;

  li {
    display: block;
    position: relative;
    width: 150px;
    aspect-ratio: 1;
    background-color: white;
    transition: background-color 0.1s ease;

    &:hover {
      background-color: rgba(var(--theme-color-rgb) / 0.7);
    }

    img {
      display: block;
      height: 100%;
      width: 100%;
      object-fit: cover;
      mix-blend-mode: multiply;
    }
  }
}

#imageViewer {
  background-color: rgba(0, 0, 0, 0.9);
  position: fixed;
  inset: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;

  .actions {
    position: absolute;
    top: 0;
  }

  img {
    max-width: 90%;
  }
}
</style>
