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
    margin-top: 10px;
    margin-left: auto;
    margin-right: auto;
    padding: 0 20px;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.home-page-button:hover {
    background-color: rgba(0,0,0,0.15);
}

.home-page-button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>