<template>
    <div class="layout">
        <Header></Header>
        <Body></Body>
        <Footer></Footer>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Header from "../header/Header.vue"
import Body from "../body/Body.vue"
import Footer from "../Footer.vue"

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "header"
        "body"
        "footer";
    grid-template-rows: 50px 1fr 50px;
}
</style>