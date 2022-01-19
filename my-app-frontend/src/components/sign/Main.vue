<template>
    <div class="main">
        <Form></Form>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Form from "./Form.vue";

const name = ref<String>("Main")
</script>

<style scoped>
.main {
    height: 100%;
    width: 100%;
    grid-area: main;
}
</style>