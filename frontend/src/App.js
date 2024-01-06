// src/App.js
import React from 'react';
import PersonForm from './components/PersonForm';
import PeopleList from './components/PeopleList';

function App() {
  const handleFormSubmit = (person) => {
    fetch('/people', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(person),
    });
  };

  return (
    <div>
      <PersonForm onFormSubmit={handleFormSubmit} />
      <PeopleList />
    </div>
  );
}

export default App;
