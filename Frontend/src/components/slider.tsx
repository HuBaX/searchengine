import * as React from 'react';
import {Box, Slider, Typography} from '@mui/material';

const marks = [
  {
    value: 0,
    label: '0 min',
  },
  {
    value: 60,
    label: '60 min',
  },
  {
    value: 120,
    label: '120 min',
  },
  {
    value: 180,
    label: '180 min',
  },
  {
    value: 240,
    label: '240 min',
  },
  {
    value: 300,
    label: '300 min',
  },
];

function valuetext(value: number) {
  return `${value}min`;
}

function DiscreteSliderLabel() {
  return (
    <Box sx={{ margin:3}} maxWidth="sm">
      <Typography variant="subtitle2">Preparation Time</Typography>
      <Slider
        defaultValue={150}
        getAriaValueText={valuetext}
        step={1}
        marks={marks}
        valueLabelDisplay="auto"
        min={0}
        max={300}
        size="medium"
      />
    </Box>
  );
}

export default DiscreteSliderLabel