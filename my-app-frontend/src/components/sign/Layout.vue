<template>
    <div class="layout">
        <PH1></PH1>
        <PH2></PH2>
        <Form></Form>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue";
import PH2 from "../placeholder/PH2.vue"
import Form from "./Form.vue";

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    height: 100%;
    width: 100%;
    background-size: cover;
    display: grid;
    grid-template-areas:
        "ph1 ph1"
        "ph2 form";
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr 3fr;
}
</style>