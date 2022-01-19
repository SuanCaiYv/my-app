<template>
    <div class="layout">
        <PH1></PH1>
        <Main></Main>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../../placeholder/PH1.vue";
import Main from "../Main.vue";

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-areas: "ph1 main";
    grid-template-columns: 1fr 1fr;
}
</style>