import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Jumbotron from 'react-bootstrap/Jumbotron';
import Form from 'react-bootstrap/Form'
import Badge from 'react-bootstrap/Badge'
import { Component } from 'react'

export default class App extends Component {
  render() {
    return (
      <Container>
        <Row><Col><Title></Title></Col></Row>
        <Row>
          <Col>
            <FibonacciInput />
            <PrimeInput />
          </Col>
        </Row>
      </Container>
    );
  }
}

function Title() {
  return (
    <div>
      <Jumbotron>
        <h1>Fibonacci and Prime numbers</h1>
        <p>
          This site lets you check whether a number is prime or Fibonacci
      </p>
      </Jumbotron>
    </div>
  );
}


class PrimeInput extends Component {
  state = {
    primeResponse: "",
    primeNumber: ""
  };

  checkNumber = () => {
    let number = parseInt(this.state.primeNumber)
    if (isNaN(number)) {
      this.setState({ primeResponse: 'False' })
      return
    }
    checkPrime(number)
      .then(res => res.json())
      .then(result =>
        this.setState({
          primeResponse: result.message
        })
      )
  };

  render() {
    return (
      <Form.Group>
        <Form.Label>Enter a number to check if it is prime</Form.Label>
        <Form.Row>
          <Col>
            <Form.Control type="number" placeholder="Number"
              onChange={(e) => {
                e.preventDefault()
                this.setState({ primeNumber: e.target.value }, () => this.checkNumber())
              }} value={this.state.primeNumber} />
          </Col>
          <Col>
            {statusIndicator(this.state.primeResponse)}
          </Col>
        </Form.Row>
      </Form.Group>
    );
  }
}

class FibonacciInput extends Component {
  state = {
    fibonacciResponse: "",
    fibonacciNumber: "",
    previous: '',
    next: ''
  };

  checkNumber = () => {
    let number = parseInt(this.state.fibonacciNumber)
    if (isNaN(number)) {
      this.setState({
        fibonacciResponse: 'False',
        previous: '',
        next: ''
      })
      return
    }
    checkFibonacci(number)
      .then(res => res.json())
      .then(result =>
        this.setState({
          fibonacciResponse: result.is_fibonacci,
          previous: result.previous,
          next: result.next
        })
      )
  };

  render() {
    return (
      <Form.Group>
        <Form.Label>Enter a number to check if it is Fibonacci</Form.Label>
        <Form.Row>
          <Col>
            <Form.Control type="number" placeholder="Number"
              onChange={(e) => {
                e.preventDefault()
                this.setState({ fibonacciNumber: e.target.value }, () => this.checkNumber())
              }} value={this.state.fibonacciNumber} />
          </Col>
          <Col>
            {statusIndicator(this.state.fibonacciResponse)}
          </Col>
          <Col>{surroundingNumbers(this.state.previous, this.state.next)}</Col>
        </Form.Row>
      </Form.Group>
    );
  }
}

function surroundingNumbers(previous, next) {
  return (
    <Button variant="primary" disabled={true}>
      <h4>
        Previous <Badge variant="secondary">{previous}</Badge>
      </h4>
      <h4>
        Next <Badge variant="secondary">{next}</Badge>
      </h4>
    </Button>
  )
}

function statusIndicator(status) {
  if (status === "True") return (<Button disabled={true} variant="success">Yes</Button>)
  if (status === "False") return (<Button disabled={true} variant="danger">No</Button>)
  return (<Button disabled={true} variant="secondary">...</Button>)
}

function checkFibonacci(number) {
  return fetch("/check_fibonacci/" + number)
}

function checkPrime(number) {
  return fetch("/check_number/" + number)
}

