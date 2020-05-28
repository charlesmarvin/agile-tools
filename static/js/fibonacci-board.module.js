const SELECTED_VOTE_BUTTON_CLASS = 'button-primary'

class Board {
  constructor () {
    this.board = null
    this.buttons = []
    this.selectedVote = null
  }

  init (targetId) {
    console.log('Initializing board in target: ' + targetId)
    this.board = document.createElement('div')
    this.board.appendChild(this._createVotingButtons())
    document.getElementById(targetId).appendChild(this.board)
  }

  _createVotingButton (text, clickHandler) {
    const button = document.createElement('button')
    button.id = `vote-${text}`
    button.innerHTML = text
    button.addEventListener('click', () => clickHandler(text))
    return button
  }

  _updateSelection (value) {
    console.log('Selected ' + value)
    this.buttons.forEach(button => {
      if (button.id === `vote-${value}`) {
        button.classList.add(SELECTED_VOTE_BUTTON_CLASS)
        this.selectedVote = button
      } else {
        button.classList.remove(SELECTED_VOTE_BUTTON_CLASS)
      }
    })
  }

  _createVotingButtons () {
    const buttonContainer = document.createElement('div')
    this.buttons = this._getButtons.map(item => this._createVotingButton(item, this._updateSelection))
    this.buttons.forEach(element => {
      buttonContainer.appendChild(element)
    })
    return buttonContainer
  }
}

export class Standard extends Board {
  _getButtons () {
    return [ 0, '1/2', 1, 2, 3, 5, 8, 13, 20, 40, 100, '?' ]
  }
}

export class Fibonacci extends Board {
  _getButtons () {
    return [ 0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, '?' ]
  }
}

export class TShirt extends Board {
  _getButtons () {
    return [ 'XS', 'S', 'M', 'L', 'XL', 'XXL', '?' ]
  }
}
