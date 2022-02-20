<template>
  <h1>{{ $t('home.about.title') }}</h1>
  <section id="about">
    <div class="row">
      <div class="secondary">
        <h4>Was</h4>
      </div>
      <div class="primary"><p>{{ $t('home.about.content.1') }}</p></div>
    </div>
    <div class="row">
      <div class="secondary">
        <h4>Wann</h4>
      </div>
      <div class="primary"><p>{{ $t('home.about.content.2') }}</p></div>
    </div>
    <div class="row">
      <div class="secondary">
        <h4>Wo</h4>
      </div>
      <div class="primary"><p>{{ $t('home.about.content.3') }}</p></div>
    </div>
    <div class="row">
      <div class="secondary"><h4>Wieso</h4></div>
      <div class="primary"><p>{{ $t('home.about.content.4') }}</p></div>
    </div>
  </section>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
.row {
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: row;
  align-items: center;

  .secondary, .primary{
    flex: 1;
  }

  &:nth-child(odd) {
    flex-direction: row-reverse;

    .secondary{
      justify-content: start;
    }
  }
}

p {
  padding: 2em;
}

.secondary{
  display: flex;
  justify-content: end;
  align-items: center;
  h4 {
    text-transform: uppercase;
    margin-inline: min(100px, 25%);
    color: white;
    border-radius: 50px;
    width: 100px;
    height: 100px;
    line-height: 100px;
    box-sizing: border-box;
    aspect-ratio: 1;
    background-color: var(--theme-color)
  }
}


</style>
