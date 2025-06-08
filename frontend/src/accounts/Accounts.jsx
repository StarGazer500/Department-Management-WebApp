
import { useState,useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';


export const CreateUserAccount = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    first_name:'',
    last_name:'',
    role_name:'',
  });
  const [error, setError] = useState(null);
  const [data, setData] = useState(null);

  const navigate = useNavigate();
  const location = useLocation();


 
  const from = location.state?.from?.pathname || "/";
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null); // Clear previous errors
    try {
      const response = await fetch(
        'http://127.0.0.1:8000/account/create-user',
        {
          method: 'POST',
          credentials: 'include', 
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            email: formData.email, 
            password: formData.password,
            first_name: formData.first_name,
            last_name: formData.last_name,
            role_name: formData.role_name,
  
          })
        }
      );

      // Log the raw response for debugging
      const rawResponse = await response.text();
      console.log('Raw response:', rawResponse);

      if (!response.ok) {
        let errorData;
        try {
          errorData = JSON.parse(rawResponse);
        } catch {
          errorData = { error: `Server returned status ${response.status}` };
        }
        console.log('User Creation Failed failed:', response.status, errorData.error || 'Unknown error');
        setError(errorData.error || 'Login failed');
        return null;
      }

      // Parse JSON only if response is OK
      const data = JSON.parse(rawResponse);
      const origin = location.state?.from || '/';
      // navigate(origin);
      setError(null);
      console.log('User created successfully:', data);
      return data;
    } catch (error) {
      console.error('Fetch Error:', error.message);
      setError(`Network error: ${error.message}`);
      return null;
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="email">email:</label>
        <input
          type="email"
          id="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="password">password:</label>
        <input
          type="password"
          id="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          required
        />
      </div>
       <div>
        <label htmlFor="first_name">first_name:</label>
        <input
          type="text"
          id="first_name"
          name="first_name"
          value={formData.first_name}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="last_name">last_name:</label>
        <input
          type="text"
          id="last_name"
          name="last_name"
          value={formData.last_name}
          onChange={handleChange}
          required
        />
      </div>
       <div>
        <label htmlFor="role_name">role_name</label>
        <input
          type="text"
          id="role_name"
          name="role_name"
          value={formData.role_name}
          onChange={handleChange}
          required
        />
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <button type="submit">Create</button>
    </form>
  );
};


export const LoginForm = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    role:''
  });
  const [error, setError] = useState(null);
  const [data, setData] = useState(null);

  const navigate = useNavigate();
  const location = useLocation();


 

  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null); // Clear previous errors
    try {
      const response = await fetch(
        'http://127.0.0.1:8000/account/login',
        {
          method: 'POST',
          credentials: 'include', 
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            email: formData.email, 
            password: formData.password,
            role:formData.role
  
          })
        }
      );

      // Log the raw response for debugging
      const rawResponse = await response.text();
      console.log('Raw response:', rawResponse);

      if (!response.ok) {
        let errorData;
        try {
          errorData = JSON.parse(rawResponse);
        } catch {
          errorData = { error: `Server returned status ${response.status}` };
        }
        console.log('Login failed:', response.status, errorData.error || 'Unknown error');
        setError(errorData.error || 'Login failed');
        return null;
      }

      // Parse JSON only if response is OK
      const data = JSON.parse(rawResponse);
    
      navigate('/is-user-valid');
      setError(null);
      console.log('route reached successfully:', data);
      return data;
    } catch (error) {
      console.error('Fetch Error:', error.message);
      setError(`Network error: ${error.message}`);
      return null;
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="email">Email:</label>
        <input
          type="email"
          id="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="password">Password:</label>
        <input
          type="password"
          id="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="role">role:</label>
        <input
          type="text"
          id="role"
          name="role"
          value={formData.role}
          onChange={handleChange}
          required
        />
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <button type="submit">Login</button>
    </form>
  );
};


export const CreateRole = () => {
  const [formData, setFormData] = useState({
    name: '',
    description: '',
  });
  const [error, setError] = useState(null);
  const [data, setData] = useState(null);

  const navigate = useNavigate();
  const location = useLocation();


 
  const from = location.state?.from?.pathname || "/";
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null); // Clear previous errors
    try {
      const response = await fetch(
        'http://127.0.0.1:8000/account/create-role',
        {
          method: 'POST',
          credentials: 'include', 
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            name: formData.name, 
            description: formData.description,
  
          })
        }
      );

      // Log the raw response for debugging
      const rawResponse = await response.text();
      console.log('Raw response:', rawResponse);

      if (!response.ok) {
        let errorData;
        try {
          errorData = JSON.parse(rawResponse);
        } catch {
          errorData = { error: `Server returned status ${response.status}` };
        }
        console.log('Role Creation Failed failed:', response.status, errorData.error || 'Unknown error');
        setError(errorData.error || 'Login failed');
        return null;
      }

      // Parse JSON only if response is OK
      const data = JSON.parse(rawResponse);
      const origin = location.state?.from || '/';
      // navigate(origin);
      setError(null);
      console.log('role created successfully:', data);
      return data;
    } catch (error) {
      console.error('Fetch Error:', error.message);
      setError(`Network error: ${error.message}`);
      return null;
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="name">name:</label>
        <input
          type="text"
          id="name"
          name="name"
          value={formData.name}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="description">description:</label>
        <input
          type="text"
          id="description"
          name="description"
          value={formData.description}
          onChange={handleChange}
          required
        />
      </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <button type="submit">Create</button>
    </form>
  );
};

// export default CreateRole

// export default LoginForm

// export default  CreateUserAccount
