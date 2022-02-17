<template>
    <div>
        <img class="img" :src="url" alt="no-img-found">
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Img")

const props = defineProps({
    url: String
})
</script>

<style scoped>
.img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}
</style>