export default {
  name: 'Card',
  props: {
    label: [String, Number],
    selected: Boolean,
    votes: [Number]
  },
  filters: {
    fmtVotes: function (value) {
      if (value === 1) {
        return `${value} vote`
      } else if (value > 1) {
        return `${value} votes`
      } else {
        return ''
      }
    }
  },
  template: `
  <div class="Card" 
    v-bind:class="{ 'Card--selected': selected }"
    >
      <label class="Card__card-label">{{ label }}</label>
      <label class="Card__vote-label" v-if="votes">{{ votes | fmtVotes }}</label>
    </div>
  `
}
