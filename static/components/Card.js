export default {
  name: 'Card',
  props: {
    label: [String, Number],
    selected: Boolean,
    votes: [Number]
  },
  template: `
  <div class="Card" 
    v-bind:class="{ 'Card--selected': selected }"
    >
      {{ label }}
      <span v-if="votes">{{ '*'.repeat(votes) }}</span>
    </div>
  `
}
