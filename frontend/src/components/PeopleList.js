// src/components/PeopleList.js
import React, { useState, useEffect } from 'react';

const PeopleList = () => {
  const [people, setPeople] = useState([]);

  useEffect(() => {
    fetch('/people')
      .then((response) => response.json())
      .then((data) => setPeople(data));
  }, []);

  return (
    <div>
      <h2>People List</h2>
      <ul>
        {people.map((person) => (
          <li key={person.email}>{`${person.name} - ${person.email}`}</li>
        ))}
      </ul>
    </div>
  );
};

export default PeopleList;
