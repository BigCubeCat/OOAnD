import * as React from "react";
import {Box, IconButton, Typography} from "@mui/joy";

import GroupsIcon from '@mui/icons-material/Groups';
import ListIcon from '@mui/icons-material/List';
import PersonOutlineIcon from "@mui/icons-material/PersonOutline";
import HomeIcon from '@mui/icons-material/Home';

const toolbarOptions = [
    {icon: <HomeIcon />, title: "Главная"},
    {icon: <GroupsIcon />, title: "Группы"},
    {icon: <ListIcon />, title: "Счета"},
    {icon: <PersonOutlineIcon />, title: "Профиль"}
];

export default function Footer(props: {}) {
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
                        sx={{display: 'flex', flexDirection:'column', alignItems: 'center'}}
                    >
                    <IconButton>{option.icon}</IconButton>
                    <Typography >{option.title}</Typography>
                    </Box>
                ))}
            </Box>
        </Box>
    );

}