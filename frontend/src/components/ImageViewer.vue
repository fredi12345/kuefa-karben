<template>
  <div id="imageViewer"
       v-if="images && images.length > 0"
       :class="{
         visible: viewerVisible,
         appear: appearViewer,
         hide: hideViewer,
         activeDrag: isDragging
       }"
       @click="closeViewer">
    <div class="actions" @click.stop>
      <button @click="closeViewer">Close</button>
      <button @click="changeImage(-1)">Previous</button>
      <button @click="changeImage(1)">Next</button>
    </div>
    <img id="bigImg" :class="{activeDrag: isDragging}" :src="images[bigImgIndex]" draggable="false"
         @click.stop/>
  </div>
</template>
<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {router} from "../plugins/routes";
import {useRoute} from "vue-router";

const props = defineProps({
  images: Array
})

const viewerVisible = ref(false);
const bigImgIndex = ref(0);
const isDragging = ref(false);

const appearViewer = ref(false);
const hideViewer = ref(false);

const route = useRoute();

function viewImage(index: number) {
  console.log("viewImage: " + index);
  bigImgIndex.value = index;
  viewerVisible.value = true;
  appearViewer.value = true;

  //TODO darf man so machen?
  router.replace({path: route.path, query: {img: bigImgIndex.value.toString()}})

  document.getElementById("bigImg").addEventListener("animationend", () => {
    appearViewer.value = false;
  })
}

/**
 *
 * @param to use -1 for previous image and 1 for next image
 */
function changeImage(to: number) {
  bigImgIndex.value += to;
  //TODO animate/slide when using touch
  //TODO show loading spinner
  if (bigImgIndex.value < 0) bigImgIndex.value = props.images.length - 1;
  if (bigImgIndex.value >= props.images.length) bigImgIndex.value = 0;

  //TODO darf man so machen?
  router.replace({path: route.path, query: {img: bigImgIndex.value.toString()}})
}

function closeViewer() {
  if (!isDragging.value) {
    hideViewer.value = true;
    const bigImg = document.getElementById("bigImg");
    bigImg.addEventListener("animationend", () => {
      viewerVisible.value = false;
      router.replace(route.path);
      hideViewer.value = false;
    }, {once: true})
  }
}

onMounted(() => {
  if (route.query.img && route.query.img !== "") {
    viewImage(Number(route.query.img));
  }
  let xStart: number;
  let xDistance: number;
  let img = document.getElementById("bigImg");

  if (img) {
    //TODO change to touchdown/touchup
    img.addEventListener("mousedown", (e) => {
      isDragging.value = true;
      xStart = e.pageX;
      xDistance = 0;
      console.log("mousedown on img");
    })
    document.documentElement.addEventListener("mousemove", (e) => {
      if (isDragging.value) {
        xDistance = e.pageX - xStart;
        if (isDragging.value) img.style.transform = `translateX(${xDistance}px)`;
      }
    })
    document.documentElement.addEventListener("mouseup", (e) => {
      if (isDragging.value) {
        isDragging.value = false;
        img.style.transform = '';
        console.log("isdragging: ", xDistance);

        //TODO close viewer when swiping down
        if (xDistance > 200) {
          changeImage(-1)
        }
        if (xDistance < -200) {
          changeImage(1)
        }
      }
    })

    document.addEventListener('keydown', (e) => {
      if (viewerVisible.value) {
        switch (e.key) {
          case "ArrowRight":
            e.preventDefault()
            changeImage(1);
            break;
          case "ArrowLeft":
            e.preventDefault()
            changeImage(-1);
            break;
          case "Escape":
            e.preventDefault()
            closeViewer();
            break;
        }
      }
    });
  }
})
defineExpose({
  viewImage
})
</script>
<style scoped lang="scss">
#imageViewer {
  --fadeDuration: 300ms;
  background-color: rgba(0, 0, 0, 0.9);
  position: fixed;
  inset: 0;
  display: none;
  justify-content: center;
  align-items: center;
  z-index: 10;

  .actions {
    position: absolute;
    top: 0;
  }

  &.activeDrag {
    pointer-events: none;
  }

  img {
    transition: transform 0.3s ease;
    max-width: 90%;
    pointer-events: initial;

    &.activeDrag {
      transition: transform 0s;
    }
  }

  &.visible {
    display: flex;
  }

  &.appear {
    animation: appear var(--fadeDuration);

    img {
      animation: grow var(--fadeDuration);
    }
  }

  &.hide {
    animation: appear var(--fadeDuration) reverse;

    img {
      animation: grow var(--fadeDuration) reverse;
    }
  }
}


@keyframes appear {
  from {
    opacity: 0
  }
  to {
    opacity: 1
  }
}

@keyframes grow {
  from {
    transform: scale(0.7)
  }
  to {
    transform: scale(1)
  }
}
</style>
