// Should be established with respect to the sequence diagram specifications
let UserID = -1
let FriendID = -1
const UserName = "name"
const FriendName = "friend"
const messagecontent = "Test Message Cypress"

function signup(name) {
    cy.visit('http://localhost:3000')
    cy.get('[value="sign_up"]')
        .click()
    cy.get('[id="namefield"]')
        .clear()
        .type(name)
    cy.get('[id="signup_submit"]')
        .click()
}

function signin(id) {
    cy.visit('http://localhost:3000')
    cy.get('[value="sign_in"]')
        .click()
    cy.get('[id="idfield"]')
        .clear()
        .type(id)
    cy.get('[id="signin_submit"]')
        .click()
}

describe('Login tests', () => {
  it('LT1: Signing up as users "name" and saving newly created unique ID', () => {
    signup(UserName)
    cy.get('[id="idPanel"]')
        .should('have.attr', 'uniqueid')
        .then(uniqueid => {
          UserID = uniqueid
        })
  })
  it('LT2: Signing in with ID stored in UserID', () => {
    signin(UserID)
  })

})

describe("Chatting frame tests: ", () => {
  // Chatting frame

  it('C1.1: Checking for users list (Registering friend)', () => {
    signup(FriendName)
    cy.get('[id="idPanel"]')
        .should('have.attr', 'uniqueid')
        .then(uniqueid => {
          FriendID = uniqueid
        })
  })
  it('C1.2: Checking for users list (Selecting friend)', () => {
    signin(UserID)
    cy.get('[value="'+FriendID+'"]')
        .click()
  })

  it('C2: Writing down messages', () => {
    signin(UserID)
    cy.get('[value="'+FriendID+'"]')
        .click()
    cy.get('textarea')
        .type(messagecontent)
    cy.get('[id="sendbutton"]')
        .click()
  })
  

})