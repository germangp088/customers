import React from 'react';
import { post } from "../request/request";

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: '',
      lastname: '',
      errorMessage: ''
    }
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChangeFirstName = this.handleChangeFirstName.bind(this);
    this.handleChangeLastName = this.handleChangeLastName.bind(this);
  }

  handleSubmit = async (e) => {
    e.preventDefault();
    let customer = {
      firstname: this.state.firstname,
      lastname: this.state.lastname
    };

    try {
      customer.id = await post(customer);
      this.props.updateCustomers(customer);
    } catch (error) {
      console.log({error})
      this.setState({ errorMessage: error.message});
    }
  }

  handleChangeFirstName(e){
    this.setState({ firstname: e.target.value});
  }

  handleChangeLastName(e){
    this.setState({ lastname: e.target.value});
  }

  render() {
    return (
      <div>
        <label>{this.state.errorMessage}</label>
        <form onSubmit={this.handleSubmit}>
          <label>First Name:</label><input type="text" value={this.state.firstname} onChange={this.handleChangeFirstName}></input><br />
          <label>Last Name:</label><input type="text" value={this.state.lastname} onChange={this.handleChangeLastName}></input><br />
          <button type="submit">Save</button>
        </form>
      </div>
    );
  }
}

export default Form;
