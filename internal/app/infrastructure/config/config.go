/*
Configuration file for Api Config
This file contains the structure of the configuration file
*/
package config

import "time"

// Worker represents the YAML structure
type Config struct {
	App struct {
		Service struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"service"`
		Template struct {
			Path string `yaml:"path"`
		} `yaml:"template"`
		Requester struct {
			Http struct {
				Timeout     time.Duration `yaml:"timeout"`
				DialContext struct {
					Timeout   time.Duration `yaml:"timeout"`
					KeepAlive time.Duration `yaml:"host"`
				} `yaml:"dial_context"`
			} `yaml:"http"`
		} `yaml:"requester"`
	} `yaml:"app"`
}
