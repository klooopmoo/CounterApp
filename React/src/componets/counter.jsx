import React, { Component } from "react";

//state is data that is local or private from the component
//props is data that we give to a component
class Counter extends Component {
  styles = {
    fontSize: 30,
    fontWeight: "bold",
  };

  //Binding since functions may not have accesss to other values
  //   constructor() {
  //     super();
  //     this.handleIncrease = this.handleIncrease.bind(this);
  //   }

  //alternative way to bind event handlers
  //   handleIncrease = () => {
  //     this.setState({ value: this.state.value + 1 });
  //     console.log("Increased clicked", this.state.value);
  //   };

  //   renderTags() {
  //     if (this.state.tags.length === 0) return <p>There are no tags</p>;
  //     return (
  //       <ul>
  //         {this.state.tags.map((tag) => (
  //           <li key={tag}>{tag}</li>
  //         ))}
  //       </ul>
  //     );
  //   }

  render() {
    return (
      <div>
        <span style={this.styles} className={this.getBadgeClasses()}>
          {this.formatCount()}
        </span>
        <button
          onClick={() => this.props.onIncrease(this.props.counter)}
          className="btn btn-secondary btn-sm"
        >
          Increase
        </button>
        <button
          onClick={() => this.props.onDelete(this.props.counter.id)}
          className="btn btn-danger btn-sm m-2"
        >
          Delete
        </button>
      </div>
    );
  }

  getBadgeClasses() {
    let classes = "badge m-2 badge-";
    classes += this.props.counter.value === 0 ? "warning" : "primary";
    return classes;
  }

  formatCount() {
    const { value: count } = this.props.counter;
    return count === 0 ? "Zero" : count;
  }
}

export default Counter;
