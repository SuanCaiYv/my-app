<template>
    <div class="layout">
        <Header></Header>
        <PH1></PH1>
        <User></User>
        <Info></Info>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Header from "./Header.vue"
import PH1 from "../placeholder/PH1.vue"
import User from "./User.vue";
import Info from "./Info.vue"

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "ph1"
        "user"
        "info";
    grid-template-rows: 60px 300px 1fr;
}
</style>