<template>
    <div class="tag">
        {{value}}
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Tag")

const props = defineProps({
    value: String
})
</script>

<style scoped>
.tag {
    width: 160px;
    height: 40px;
    margin: 0 20px 10px;
    padding: 0;
    border: 1px solid black;
    box-sizing: border-box;
    border-radius: 6px;
    line-height: 40px;
    font-size: 1.2rem;
}
</style>