<template>
  <section>
    <main>
      <h1>{{ event.title }}</h1>
      <img :src="event.img"/>

      <div id="menu">
        <p class="menuItem menu__1">{{event.vorspeise}}</p>
        <p class="menuItem menu__2">{{event.hauptgang}}</p>
        <p class="menuItem menu__3">{{event.nachtisch}}</p>
      </div>
      <p>{{ event.description }}</p>
    </main>
    <aside>
      <div id="registration">
        <h2>Anmeldung</h2>
        <p class="info closing">Anmeldeschluss: {{ event.closingDate}}</p>
        <form>
          <input class="name" type="text" placeholder="Name" required><br>
          <textarea rows="2" name="message" id="participationMessage" type="text"
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
      <ol>
        <li v-for="comment in event.comments">
        <span>{{comment.name}} </span>
        <span> {{ comment.date}}</span>
          <p>{{comment.text}}</p>
        </li>
      </ol>
    </div>
    </aside>
  </section>
</template>

<script setup lang="ts">
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
    {name: "Si.",
    date: "10.03.2019",
    text: "Ich möchte auch gern zum Krimidinner kommen. Wann gibt es so etwas nochmal?"}

  ]
}
</script>

<style scoped lang="scss">
section{
  display: flex;
  flex-direction: row;
  gap: 40px;
  padding-inline: 40px;

}
aside {
  div{
    padding-inline: 20px;
    padding-bottom: 20px;
    overflow: auto;
    background-color: var(--card-color);
    text-align: start;
    margin-top: 20px;
  }
}
#menu{
  .menuItem{
    transform: translate(-50%);
    position: relative;
  }
  .menu__1{
    left: 25%;
  }
  .menu__2{
    left: 50%;
  }
  .menu__3{
    left: 75%;
  }
}
#registration{
  input, textarea{
    margin-top: 6px;
  }
  .name, textarea{
    box-sizing: border-box;
    width: 100%;
  };
  textarea{
    resize: vertical;
    font-family: inherit;
  }

  .menuCount {
    display: block;
    padding-inline-start: 20px;
    input{
      width: 30px;
      margin-inline-end: 6px;
    }

  }
}

#comments{
  ol{
    list-style: none;
    padding: 0;
  };
}
</style>
