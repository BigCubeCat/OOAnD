import {Box, Typography} from "@mui/joy";

import GroupsIcon from '@mui/icons-material/Groups';
import ListIcon from '@mui/icons-material/List';
import PersonOutlineIcon from "@mui/icons-material/PersonOutline";
import HomeIcon from '@mui/icons-material/Home';
import {useLocation} from "wouter";

type TToolbarOption = {
    icon: any,
    title: string,
    url: string,
};

const toolbarOptions: TToolbarOption[] = [
    {icon: <HomeIcon sx={{ fontSize: 32 }}/>, title: "Главная", url: '/'},
    {icon: <GroupsIcon sx={{ fontSize: 32 }}/>, title: "Группы", url: '/groups'},
    {icon: <ListIcon sx={{ fontSize: 32 }}/>, title: "Счета", url: '/bills'},
    {icon: <PersonOutlineIcon sx={{ fontSize: 32 }}/>, title: "Профиль", url: '/me'}
];

const locationDict: {[id: string] : number} = {
    '/': 0,
    'groups': 1,
    'bills': 2,
    'profile': 3,
};

export default function Footer() {
    const [location, navigate] = useLocation();
    const lockId = locationDict[location];
    console.log(lockId);
    return(
        <Box sx={{
            bgcolor: 'secondary.700',
            padding: '1em',
            display: 'flex',
            justifyContent: 'space-between', alignItems: 'center'
        }}>
            <Box sx={{
                display: 'flex',
                justifyContent: 'space-around',
                width: "100%"
            }}>

                {toolbarOptions.map((option, index) => (
                    <Box
                        key={"menu-"+index}
                        sx={{
                            display: 'flex',
                            flexDirection:'column',
                            alignItems: 'center',
                        }}
                        onClick={() => navigate(option.url)}
                    >
                    {option.icon}
                    <Typography>{option.title}</Typography>
                    </Box>
                ))}
            </Box>
        </Box>
    );

}