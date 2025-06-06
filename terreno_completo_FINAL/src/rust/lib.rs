use wasm_bindgen::prelude::*;
use web_sys::{Document, Element};
use serde::Deserialize;
use serde_yaml::Value as YamlValue;
use std::collections::BTreeMap;

// Define a estrutura básica de um span ui.render
#[derive(Debug, Deserialize)]
struct Span {
    id: Option<String>,
    #[serde(rename = "type")]
    span_type: String,
    blocks: Option<Vec<Block>>,
}

#[derive(Debug, Deserialize)]
struct Block {
    component: String,
    #[serde(default)]
    props: BTreeMap<String, YamlValue>,
    #[serde(default)]
    children: Option<Vec<YamlValue>>, // simplificado, para strings ou blocos aninhados
}

#[wasm_bindgen(start)]
pub fn run() {
    // Permite erros de panic aparecerem no console do navegador
    console_error_panic_hook::set_once();
}

// Função exposta ao JS para renderizar um contrato .logline no <div id="app">
#[wasm_bindgen]
pub fn render_contract(logline_text: &str) -> Result<(), JsValue> {
    let contract: Vec<Span> = serde_yaml::from_str(logline_text)
        .map_err(|e| JsValue::from_str(&format!("YAML parse error: {}", e)))?;
    let window = web_sys::window().unwrap();
    let document: Document = window.document().unwrap();
    let app_div: Element = document.get_element_by_id("app")
        .ok_or_else(|| JsValue::from_str("Elemento #app não encontrado"))?;

    // Limpa conteúdo
    app_div.set_inner_html("");

    for span in contract {
        if span.span_type == "ui.render" {
            if let Some(blocks) = span.blocks {
                for block in blocks {
                    let child_el = create_element_from_block(&document, &block)?;
                    app_div.append_child(&child_el)?;
                }
            }
        }
    }
    Ok(())
}

// Converte um Block em um Element DOM
fn create_element_from_block(document: &Document, block: &Block) -> Result<Element, JsValue> {
    let tag = &block.component;
    let el = document.create_element(tag)
        .map_err(|e| JsValue::from_str(&format!("Falha ao criar <{}>: {:?}", tag, e)))?;

    // Atribui props (class, text, style, onclick, etc.)
    for (key, val) in &block.props {
        match key.as_str() {
            "class" => {
                if let Some(s) = val.as_str() {
                    el.set_attribute("class", s)?;
                }
            }
            "text" => {
                if let Some(s) = val.as_str() {
                    el.set_text_content(Some(s));
                }
            }
            "style" => {
                // Suporta um mapa YAML de estilos
                if let Some(style_map) = val.as_mapping() {
                    let mut style_str = String::new();
                    for (k, v) in style_map {
                        if let (Some(ks), Some(vs)) = (k.as_str(), v.as_str()) {
                            style_str.push_str(&format!("{}:{};", ks, vs));
                        }
                    }
                    el.set_attribute("style", &style_str)?;
                }
            }
            "action" => {
                // Trata ui.click: atribui evento onclick que gera span ui.click
                if let Some(action_name) = val.as_str() {
                    let action = action_name.to_string();
                    let closure = Closure::wrap(Box::new(move || {
                        // Gera um span de clique; aqui só logamos no console
                        web_sys::console::log_1(&JsValue::from_str(
                            &format!(r#"{{\"type\":\"ui.click\",\"action\":\"{}\",\"timestamp\":\"{}\"}}"#,
                                action,
                                js_sys::Date::new_0().to_iso_string()
                            )
                        ));
                    }) as Box<dyn FnMut()>);
                    el.add_event_listener_with_callback("click", closure.as_ref().unchecked_ref())?;
                    closure.forget();
                }
            }
            "href" => {
                if let Some(s) = val.as_str() {
                    el.set_attribute("href", s)?;
                }
            }
            _ => {
                // Outras props genéricas: data-*
                if let Some(s) = val.as_str() {
                    el.set_attribute(key, s)?;
                }
            }
        }
    }

    // Caso haja children aninhados (HTML bruto ou novos blocks em YAML),
    // poderíamos parsear recursivamente. Para simplicidade, ignoramos aqui.

    Ok(el)
}
