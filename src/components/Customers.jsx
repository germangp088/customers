import React from 'react';
import './Customers.css';
import { get } from "../request/request";
import {
  Accordion,
  AccordionItem,
  AccordionItemHeading,
  AccordionItemPanel,
  AccordionItemButton
} from 'react-accessible-accordion';
import Form from './Form';

class Customers extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      customers: []
    };
    this.updateCustomers = this.updateCustomers.bind(this);
  }

  async componentDidMount() {
    const customers = await get();
    this.setState({
      customers: customers || []
    });
  }

  updateCustomers(customer) {
    this.setState({
      customers: [...this.state.customers, customer]
    });
  }

  render() {
      const accordionItems = this.state.customers.map(customer => {
        const { id, firstname, lastname } = customer

        return (
            <AccordionItem key={id}>
            <AccordionItemHeading>
                <AccordionItemButton className={`accordion__button accordion__button--card`}>
                    {`First Name: ${firstname} - Last Name: ${lastname}`}
                </AccordionItemButton>
            </AccordionItemHeading>
            <AccordionItemPanel>
                <ul>
                    <li>{`Id : ${id}`}</li>
                    <li>{`First Name : ${firstname}`}</li>
                    <li>{`Last Name : ${lastname}`}</li>
                </ul>
            </AccordionItemPanel>
        </AccordionItem>
        )
    })
    return (
      <Accordion allowMultipleExpanded={true} allowZeroExpanded ={true}>
        {accordionItems.length === 0 ? <div>No customers to show.</div> : 
        <div>
          {accordionItems}
          <Form updateCustomers={this.updateCustomers} />
        </div> }
      </Accordion>
    );
  }
}

export default Customers;
