<template>
    <div class="home-page">
        <button class="home-page-button" @click="router.push('/home')">CodeWithBuff</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";

const name = ref<String>("Home")

const router = useRouter()
</script>

<style scoped>
.home-page {
    width: 100%;
    height: 100%;
    grid-area: home-page;
}

.home-page-button {
    width: auto;
    height: 40px;
    margin: 10px auto;
    padding: 0 20px;
    border: 2px solid lightgray;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
    font-weight: bolder;
    background-color: white;
}

.home-page-button:hover {
    background-color: lightgray;
}

.home-page-button:active {
    background-color: gainsboro;
}
</style>