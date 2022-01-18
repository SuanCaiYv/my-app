<template>
    <div>
        <input type="text" v-model="str">
        {{str}}
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Test")

let str = ref<String>("")
</script>

<style scoped>

</style>