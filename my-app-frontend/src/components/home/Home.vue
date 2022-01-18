<template>
    <div class="home">
        <button class="homePage" @click="homePage">
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
.home {
    width: 100%;
    height: 100%;
}

.homePage {
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    border-radius: 8px;
    border: 0;
    background-color: lightskyblue;
}

.homePage:hover {
    background-color: deepskyblue;
}

.homePage:active {
    background-color: dodgerblue;
}
</style>