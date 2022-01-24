<template>
    <div class="home-page">
        <button class="home-page-button" @click="homePage">
            <span style="font-size: 1.2rem">Home</span>
        </button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";

const name = ref<String>("Home")

const router = useRouter()
const homePage = function () {
    router.push("/home")
}
</script>

<style scoped>
.home-page {
    width: 100%;
    height: 100%;
    grid-area: home-page;
}

.home-page-button {
    width: 120px;
    height: 100%;
    margin-left: 40px;
    margin-right: 40px;
    padding: 0;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    background-color: lightskyblue;
}

.home-page-button:hover {
    background-color: deepskyblue;
}

.home-page-button:active {
    background-color: dodgerblue;
}
</style>