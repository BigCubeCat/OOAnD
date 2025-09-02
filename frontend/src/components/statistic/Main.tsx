import {useLocation} from "wouter";

export default  function Main() {
    const [location] = useLocation();
    return (
        <div>{location}</div>
    )
}