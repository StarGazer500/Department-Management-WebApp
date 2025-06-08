
import { useState,useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';


export const LecturerPage = () => {

  const [error, setError] = useState(null);


//   const navigate = useNavigate();
//   const location = useLocation();

  

   useEffect(() => {
  
    setError(null); // Clear previous errors
    const handleCheckUserValidity=async()=>{
    try {
      const response = await fetch(
        'http://127.0.0.1:8000/account/is-user-valid',
        {
          method: 'GET',
          credentials: 'include', 
          headers: { 'Content-Type': 'application/json' },
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
        console.log('User Check Failed:', response.status, errorData.error || 'Unknown error');
        setError(errorData.error || 'user check failed failed');
        return null;
      }

      // Parse JSON only if response is OK
      const data = JSON.parse(rawResponse);
     
    //   navigate(origin);
      setError(null);
      console.log('user is still in login session:', data);
      return data;
    } catch (error) {
      console.error('Fetch Error:', error.message);
      setError(`Network error: ${error.message}`);
      return null;
    }
 }

 handleCheckUserValidity()
})


  
  

  return (
    <div>
        <p>welcome to Lecuters page</p>
    </div>
  );
};


