<template>
  <section id="eventView">
    <main>
      <div v-if="auth.loggedIn">
        <button>Edit</button>
        <button>Delete</button>
      </div>

      <EventDetails :event="event"/>

    </main>
    <aside>
      <div id="registration">
        <h2>Anmeldung</h2>
        <p class="info closing">Anmeldeschluss: {{ event.closingDate }}</p>
        <form>
          <input class="name" type="text" placeholder="Name" required><br>
          <textarea rows="3" name="message" id="participationMessage"
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
        <CommentList :comments="event.comments"/>
        <form id="commentForm">
          <textarea rows="3" placeholder="Kommentar"></textarea>
          <button>Senden</button>
        </form>
      </div>
    </aside>
  </section>
</template>

<script setup lang="ts">
import EventDetails from "../components/EventDetails.vue";
import {CreateEventRequest} from "../api/generated";
import {useAuth} from "../stores/auth";
import CommentList from "./CommentList.vue";

const auth = useAuth();
const event: CreateEventRequest = {
  theme: "Heimische Kräuterküche",
  startingDate: "22.02.2022",
  img: "https://unsplash.it/640/425",
  description: "Lord Grey bittet zum Gärtnerkongress! Seinem Ruf folgen Botaniker, Lifestyle-Blogger und Imker, um sich über die neusten Erfindungen der Branche auszutauschen. Doch schon bald müssen sich die Gäste einer jahrzehntealten Frage stellen: Ist der Mörder immer der Gärtner?",
  starter: "Handkäs-Bärlauch-Dip mit geröstetem Brot",
  mainDish: "Grüne Soße mit Kartoffeln und Ei",
  dessert: "Zitronen-Basilikum-Sorbet",
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

</script>

<style scoped lang="scss">
#eventView {
  display: flex;
  flex-direction: row;
  gap: 40px;
  padding-inline: 3%;
  flex-wrap: wrap;
  padding-top: 20px;
}

main {
  flex: 4;
  flex-basis: 500px;
  padding-top: 10px;
}

aside {
  --padding: 20px;
  flex: 1;

  flex-basis: 350px;

  > div {
    padding-bottom: var(--padding);
    overflow: auto;
    background-color: var(--card-color);
    text-align: start;
    &+div{
      margin-top: var(--padding);
    }
  }

  input:not(.menuItem), textarea {
    margin-top: 6px;
    box-sizing: border-box;
    width: 100%;
  }

  button {
    float: right;
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

  #commentForm {
    textarea {
      box-sizing: border-box;
      width: 100%;
    }
  }
}

</style>
